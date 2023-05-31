.DEFAULT_GOAL    := help

## ###############################################################
## Parameters and setups
## #################################################################
IMPORT_PATH      := github.com/KvalitetsIT/kih-telecare-exporter
DOCKER_IMAGE     := oth/exporter
ECR_REPO         := 401334847138.dkr.ecr.eu-west-1.amazonaws.com/${DOCKER_IMAGE}
VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H:%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.Build=${DATE}"'

BRANCH := "master"
TAG ?=

dockertest       :=

ifndef TAG
DOCKER_TAG := master
TAG := master
else
DOCKER_TAG := ${TAG}
endif


build_dir        := $(CURDIR)/build
dist_dir         := $(CURDIR)/dist
exec             := $(DOCKER_IMAGE)

# assuming go 1.9 here!!
_allpackages = $(shell go list ./...| grep -v '/vendor/' )

# memoize allpackages, so that it's executed only once and only if used
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)

# Figure out, loud or quiet?
Q := $(if $V,,@)


# comment this line out for quieter things
#V?=? # When V is set, print commands and build progress.
VERBOSE?=false
# Space separated patterns of packages to skip in list, test, format.
PRJ_SRCS=$$(go list ./... | grep -v '/vendor/' | grep -v '/builtin/bins/')

ifeq ($(VERBOSE),true)
	V:=1
endif

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

## #################################################################
## OTH Targets
## #################################################################
start: ## Starts in development mode
	@echo "Starting development service"
	reflex -v -s -- bash -c 'make serve'

## #################################################################
## Local targets
## #################################################################

resetdb: build ### Drop current database and re-apply migrations
	exporter migrate drop
	exporter migrate

mysql: ### connnect to dev db
	docker exec -it devdb bash -c 'mysql -uroot -p$$MYSQL_ROOT_PASSWORD exporter'

serve: test build ### Just starts API
	$(build_dir)/exporter serve


# safebuild builds inside a docker container with no clingons from your $GOPATH
safebuild:
	@echo "Building..."
	mkdir -p build
	$Q docker build -t $(DOCKER_IMAGE):$(VERSION) .

build: ###
	@echo "Building..."
	$Q go build -o $(build_dir)/exporter $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)

tags: ### Displays tags
	@echo "Listing tags..."
	$Q @git tag

tag: ## Set version
	@echo "tagging repo... "
ifndef TAG
	$(error "TAG is not - please set it")
else
	@echo "Version is: $$TAG"
	@echo "Tagging version..."
	git tag -a $$TAG -m "Tagging $$TAG" -f
	@git push --tags
endif

buildcontainer: ### Builds docker container
	@docker build          \
	--build-arg version=${DOCKER_TAG}       \
	-t ${DOCKER_IMAGE}:${DOCKER_TAG} .

ecr-login: ### Performs ECR login
	@aws ecr get-login-password \
    --region eu-west-1 \
	| docker login \
    --username AWS \
    --password-stdin 401334847138.dkr.ecr.eu-west-1.amazonaws.com

tag-container: ### tags docker image
	@echo "tagging - ${DOCKER_TAG}"
	docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${ECR_REPO}:${DOCKER_TAG}
	@echo "Done tagging"

docker-logs: ### Tails logs from container
	docker exec exporter tail -f /var/log/exporter/stdout/current

docker-run: buildcontainer ### buildcontainer application in container
	docker run --rm -d --name exporter --network openteledev \
	-v $$(pwd)/docker/conf/exporter.yaml:/app/exporter.yaml:ro \
	-p 8360:8360 \
	${DOCKER_IMAGE}:${DOCKER_TAG}

docker-stop: ### Stop running container
	@docker stop exporter

docker-enter: ### Enter container
	@docker exec -it exporter bash

push-to-ecr: ### Push container to ecr
	@echo "Pushing - ${ECR_REPO}:${DOCKER_TAG}"
	@docker push ${ECR_REPO}:${DOCKER_TAG}

release: clean ecr-login buildcontainer tag-container push-to-ecr ## Release component
	@echo "Built docker container and pushed to AWS... $(TAG)"
	@if [ ! -d release ]; then mkdir release; fi
ifeq ($(TAG), master)
	git archive --format zip --output release/exporter.zip origin/master
else
	git archive --format zip --output release/exporter.zip $(TAG)
endif

dockerize: ## Dockerize component
	@echo "Docker image already build and pushed as part of release target"

### Code not in the repository root? Another binary? Add to the path like this.
# .PHONY: otherbin
# otherbin: .GOPATH/.ok
#   $Q go install $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/otherbin

##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Utility targets <===== #####

deps: setup
	@echo "Ensuring Dependencies..."
	$Q go env

docker:
	@echo "Docker Build..."
	$Q docker build -t $(DOCKER_IMAGE) .

clean:  ## Cleans source and dependencies
	@echo "Clean..."
	$Q rm -rf bin


cleanall: clean clean-dist clean-build clean-cover ## Cleans everything

clean-build: ### Removes croos compilation files
	@echo "Removing cross-compilation files"
	rm -rf $(build_dir)

clean-dist: ### Removes distribution files
	@echo "Removing distribution files"
	rm -rf $(dist_dir)
clean-cover: ### Removes coverage files
	@echo "Removing cover files"
	rm -rf cover

compile: build ## Compile project
	@echo "Compiling project ..."
	@go build -o exporter

test: ## Runs test for component
	@echo "Testing..."
ifndef SSH_CONNECTION
	@echo "Testing Outside CI..."
	$Q go vet $(allpackages)
	$Q GODEBUG=cgocheck=2 go test -race $(allpackages)
else
	docker run --rm -v $$(pwd):/app -w /app golang:1.16 make ci-test
endif

ci-test: ### Run tests for cmponent
	@echo "Testing in CI..."
	go get -u github.com/jstemmer/go-junit-report
#	$Q go test $(if $V,-v) -i -race $(allpackages) # install -race libs to speed up next run
	mkdir -p build/test
	go test -v ./... | go-junit-report > build/test/report.xml
	$Q mkdir -p test
	$Q ( go vet $(allpackages); echo $$? ) | \
	   tee test/vet.txt | sed '$$ d'; exit $$(tail -1 test/vet.txt)
	$Q ( GODEBUG=cgocheck=2 go test -v -race $(allpackages); echo $$? ) | \
	   tee test/output.txt | sed '$$ d'; exit $$(tail -1 test/output.txt)

integrationtest: ### Runs integration tests
	@echo "Testing Outside CI..."
	$Q go vet $(allpackages)
	$Q GODEBUG=cgocheck=2 INTEGRATION_TEST=true go test -v -race $(allpackages)

testall: test integrationtest ### Test all then things

list: ### List packages
	@echo "List..."
	@echo $(allpackages)

cover: $(GOPATH)/bin/gocovmerge ### Runs coverage
	@echo "Coverage Report..."
	@echo "NOTE: make cover does not exit 1 on failure, don't use it to check for tests success!"
	$Q mkdir -p cover
	$Q rm -f cover/*.out cover/all.merged
	$(if $V,@echo "-- go test -coverpkg=./... -coverprofile=cover/... ./...")
	@for MOD in $(allpackages); do \
		go test -coverpkg=`echo $(allpackages)|tr " " ","` \
			-coverprofile=cover/unit-`echo $$MOD|tr "/" "_"`.out \
			$$MOD 2>&1 | grep -v "no packages being tested depend on"; \
	done
	$Q $(GOPATH)/bin/gocovmerge cover/*.out > cover/all.merged
ifndef CI
	@echo "Coverage Report..."
	$Q go tool cover -html cover/all.merged
else
	@echo "Coverage Report In CI..."
	$Q go tool cover -html cover/all.merged -o cover/all.html
endif
	@echo ""
	@echo "=====> Total test coverage: <====="
	@echo ""
	$Q go tool cover -func cover/all.merged

format: $(GOPATH)/bin/goimports ## Runs formatter
	@echo "Formatting..."
	$Q find . -iname \*.go | grep -v \
		-e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)) | xargs goimports -w


##### =====> Internals <===== #####
setup: clean ## Runs setup - if any
	@echo "Setup..."
	if ! grep "/bin" .gitignore > /dev/null 2>&1; then \
		echo "/bin" >> .gitignore; \
	fi
	if ! grep "/dist" .gitignore > /dev/null 2>&1; then \
		echo "/dist" >> .gitignore; \
	fi
	if ! grep "/build" .gitignore > /dev/null 2>&1; then \
		echo "/build" >> .gitignore; \
	fi
	mkdir -p cover
	mkdir -p bin
	mkdir -p test
	go get github.com/wadey/gocovmerge
	go get golang.org/x/tools/cmd/goimports

$(GOPATH)/bin/gocovmerge:
	@echo "Checking Coverage Tool Installation..."
	$Q go install github.com/wadey/gocovmerge

$(GOPATH)/bin/goimports:
	@echo "Checking Import Tool Installation..."
	@test -d $(GOPATH)/src/golang.org/x/tools/cmd/goimports || \
		{ echo "Vendored goimports not found, try running 'make setup'..."; exit 1; }
	$Q go install golang.org/x/tools/cmd/goimports


## #################################################################
## Generate help
## #################################################################
# Handles nice readme in shell and documentation
S ?=
ifndef S
format:= "\033[36m%-20s\033[0m %s\n"
else
format:= "%-20s %s\n"
endif

help: ## This help
	@printf '=%.0s' {1..80}
	@echo -e "\nStandard OTH targets:"
	@printf '=%.0s' {1..80}
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |grep -v "###"| sort | awk 'BEGIN {FS = ":.*?## "}; {printf ${format}, $$1, $$2}'
	@echo
	@printf '=%.0s' {1..80}
	@echo -e "\nExtra targets:"
	@printf '=%.0s' {1..80}
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?### .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?### "}; {printf ${format}, $$1, $$2}'


.PHONY: setup dev serve safebuild build tags release clean test list cover format docker deps clean-build clean-cover clean-dist clean setup format  compile test start set-version tag release dockerize help bootRun buildcontainer ecr-login tag-container push-to-ecr testtarget
