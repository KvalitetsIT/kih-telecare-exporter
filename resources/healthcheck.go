/// Handles health check functionality
package resources

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// Performs health check on required resources
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheck := healthCheckResponse{}

	errors := checkHealth()

	if len(errors) > 0 {
		res := healthCheckResponse{Errors: errors}
		w.WriteHeader(http.StatusServiceUnavailable)
		render.JSON(w, r, res)
		return
	}

	healthCheck.APIVersion = config.Version
	healthCheck.Environment = config.Environment

	render.JSON(w, r, healthCheck)
}

type healthCheckResponse struct {
	Message     string `json:"message,omitempty"`
	APIVersion  string `json:"apiVersion,omitempty"`
	Environment string `json:"environment,omitempty"`

	Errors []healthCheckError `json:"errors,omitempty"`
}

// Check if exporter backend is alright
func checkExporterHealth() error {
	return exprtr.CheckHealth()
}

// Performs the actual health check
func checkHealth() []healthCheckError {
	errors := []healthCheckError{}

	// check db
	if err := repo.CheckRepository(); err != nil {
		errors = append(errors, healthCheckError{Resource: "repository", Error: fmt.Sprintf("Error testing repository - %s", err)})
	}

	if err := checkExporterHealth(); err != nil {
		errors = append(errors, healthCheckError{Resource: "exporter", Error: fmt.Sprintf("Error testing exporter - %s", err)})
	}

	return errors
}
