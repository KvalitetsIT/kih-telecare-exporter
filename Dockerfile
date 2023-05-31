FROM golang:alpine AS builder

ARG version=1.0.0
ENV VERSION=$version
WORKDIR /app/

COPY . .

RUN apk add --no-cache make git openssh-client && make build

FROM golang:alpine AS runner

ARG version=1.0.0
ENV ENV=production \
    TZ=Europe/Copenhagen \
    VERSION=$version \
    TIMEZONE=$TZ \
    LANG=C.UTF-8

RUN mkdir /app

COPY --from=builder /app/build/exporter /app/exporter

RUN apk add --no-cache curl ca-certificates nss && \
    echo "Build complete"

ADD docker/root /
COPY migrations /app/migrations
EXPOSE 8360