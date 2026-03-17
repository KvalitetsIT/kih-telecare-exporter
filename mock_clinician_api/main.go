package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/measurements", measurementsHandler)
	mux.HandleFunc("/patients/", patientHandler)
	mux.HandleFunc("/measurements/", measurementHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/clinician/api/patients/", patientHandler)

	// Simple logging middleware
	loggedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		mux.ServeHTTP(w, r)
	})

	fmt.Println("Mock KIH API server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", loggedMux))
}

func measurementsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, serverResponse)
}

func patientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract patient ID from URL
	parts := strings.Split(r.URL.Path, "/")
	patientID := parts[len(parts)-1]

	// Return a patient-specific response
	patientJSON := fmt.Sprintf(`{
		"createdDate": "2019-11-11T14:38:37",
		"uniqueId": "%s",
		"username": "NancyAnn",
		"firstName": "Nancy Ann",
		"lastName": "Berggren",
		"dateOfBirth": "25-12-1948",
		"sex": "female",
		"status": "active",
		"address": "Åbogade 15",
		"postalCode": "8200",
		"city": "Aarhus N",
		"links": {
			"self": "http://localhost:8081/clinician/api/patients/%s"
		}
	}`, patientID, patientID)

	fmt.Fprint(w, patientJSON)
}

func measurementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{
		"timestamp": "2019-11-11T14:38:48.000Z",
		"type": "weight",
		"measurement": {
			"unit": "kg",
			"value": 80
		},
		"links": {
			"measurement": "http://localhost:8081/clinician/api/patients/13/measurements/3",
			"patient": "http://localhost:8081/clinician/api/patients/13"
		}
	}`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

const serverResponse = `
{
  "results": [
    {
      "timestamp": "2019-11-11T14:38:57.000Z",
      "type": "blood_pressure",
      "measurement": {
        "unit": "mmHg",
        "systolic": 130,
        "diastolic": 65
      },
      "origin": {
        "manualMeasurement": {
          "enteredBy": "clinician"
        }
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/13/measurements/214",
        "patient": "http://localhost:8081/clinician/api/patients/13"
      }
    },
    {
      "timestamp": "2019-11-11T14:38:57.000Z",
      "type": "blood_pressure",
      "measurement": {
        "unit": "mmHg",
        "systolic": 130,
        "diastolic": 65
      },
      "origin": {
        "manualMeasurement": {
          "enteredBy": "clinician"
        }
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/14/measurements/215",
        "patient": "http://localhost:8081/clinician/api/patients/14"
      }
    },
    {
      "timestamp": "2019-11-11T14:38:48.000Z",
      "type": "weight",
      "measurement": {
        "unit": "kg",
        "value": 80
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/13/measurements/3",
        "patient": "http://localhost:8081/clinician/api/patients/13"
      }
    }
  ]
}
`
