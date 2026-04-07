package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Measurement struct {
	Timestamp   string      `json:"timestamp"`
	Type        string      `json:"type"`
	Measurement interface{} `json:"measurement"`
	Origin      *Origin     `json:"origin,omitempty"`
	Links       Links       `json:"links"`
}

type BloodPressureMeasurement struct {
	Unit      string `json:"unit"`
	Systolic  int    `json:"systolic"`
	Diastolic int    `json:"diastolic"`
}

type WeightMeasurement struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Origin struct {
	ManualMeasurement ManualMeasurement `json:"manualMeasurement"`
}

type ManualMeasurement struct {
	EnteredBy string `json:"enteredBy"`
}

type Links struct {
	Measurement string `json:"measurement"`
	Patient     string `json:"patient"`
}

type ServerResponse struct {
	Results []Measurement `json:"results"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/measurements", measurementsHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/clinician/api/", clinicianAPIHandler)

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
	now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	// now := "2026-03-19T10:05:00.000Z"
	response := fmt.Sprintf(serverResponse, now, now, now, now, now, now)
	fmt.Fprint(w, response)
}

/*
// Kode til at simulere 10.000 målinger til at teste, hvordan systemet håndterer store datamængder.
func measurementsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	measurements := make([]Measurement, 10000)
	for i := 0; i < 10000; i++ {
		patientID := rand.Intn(100) + 1 // Random patient ID between 1 and 100
		measurementID := 214 + i

		var measurementData interface{}
		var origin *Origin
		var mType string

		// Alternate between blood pressure and weight
		if i%2 == 0 {
			mType = "blood_pressure"
			measurementData = BloodPressureMeasurement{
				Unit:      "mmHg",
				Systolic:  rand.Intn(40) + 110, // 110-149
				Diastolic: rand.Intn(30) + 60,  // 60-89
			}
			origin = &Origin{
				ManualMeasurement: ManualMeasurement{
					EnteredBy: "clinician",
				},
			}
		} else {
			mType = "weight"
			measurementData = WeightMeasurement{
				Unit:  "kg",
				Value: rand.Intn(50) + 50, // 50-99
			}
			origin = nil
		}

		measurements[i] = Measurement{
			Timestamp:   time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			Type:        mType,
			Measurement: measurementData,
			Origin:      origin,
			Links: Links{
				Measurement: fmt.Sprintf("http://localhost:8081/clinician/api/patients/%d/measurements/%d", patientID, measurementID),
				Patient:     fmt.Sprintf("http://localhost:8081/clinician/api/patients/%d", patientID),
			},
		}
	}

	response := ServerResponse{Results: measurements}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}*/

// 3 eksempler fra vores testsystem
var patients = map[string]string{
	"1610940125": `{
		"createdDate": "%s",
		"uniqueId": "1610940125",
		"username": "BrianB",
		"firstName": "Brian",
		"lastName": "Bach",
		"dateOfBirth": "16-10-1994",
		"sex": "male",
		"status": "active",
		"address": "Højgårdsparken 444",
		"postalCode": "1438",
		"city": "København K",
		"links": {
			"self": "http://localhost:8081/clinician/api/patients/1610940125"
		}
	}`,
	"1805039414": `{
		"createdDate": "%s",
		"uniqueId": "1805039414",
		"username": "HelleL",
		"firstName": "Helle",
		"lastName": "Larsen",
		"dateOfBirth": "16-10-1994",
		"sex": "female",
		"status": "active",
		"address": "Rungsted Strandvej 97",
		"postalCode": "6690",
		"city": "Gørding",
		"links": {
			"self": "http://localhost:8081/clinician/api/patients/1805039414"
		}
	}`,
	"2303079354": `{
		"createdDate": "%s",
		"uniqueId": "2303079354",
		"username": "AneP",
		"firstName": "Ane",
		"lastName": "Pedersen",
		"dateOfBirth": "03-03-2000",
		"sex": "female",
		"status": "active",
		"address": "Islands Brygge 260",
		"postalCode": "3630",
		"city": "Jægerspris",
		"links": {
			"self": "http://localhost:8081/clinician/api/patients/2303079354"
		}
	}`,
}

func clinicianAPIHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	// /clinician/api/patients/{id}
	// /clinician/api/patients/{id}/measurements/{id}
	if len(parts) >= 5 && parts[3] == "patients" {
		if len(parts) >= 7 && parts[5] == "measurements" {
			measurementHandler(w, r)
			return
		}
		patientHandler(w, r)
		return
	}

	http.NotFound(w, r)
}

func patientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Extract patient ID from URL
	parts := strings.Split(r.URL.Path, "/")
	patientID := parts[4]

	// Return a patient-specific response
	now := time.Now().UTC().Format("2006-01-02T15:04:05")

	patientJSON, ok := patients[patientID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, fmt.Sprintf(patientJSON, now))
}

func measurementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	patientID := parts[4]
	measurementID := parts[6]

	// This is not efficient, but for a mock server it's fine.
	var serverResp ServerResponse
	now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	responseString := fmt.Sprintf(serverResponse, now, now, now, now, now, now)
	if err := json.Unmarshal([]byte(responseString), &serverResp); err != nil {
		http.Error(w, "Failed to parse measurements", http.StatusInternalServerError)
		return
	}

	for _, m := range serverResp.Results {
		linkParts := strings.Split(m.Links.Measurement, "/")
		if len(linkParts) >= 7 {
			pID := linkParts[len(linkParts)-3]
			mID := linkParts[len(linkParts)-1]
			if pID == patientID && mID == measurementID {
				json.NewEncoder(w).Encode(m)
				return
			}
		}
	}

	http.NotFound(w, r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var serverResponse = `
{
  "results": [
    {
      "timestamp": "%s",
      "type": "blood_pressure",
      "measurement": {
        "unit": "mmHg",
        "systolic": 120,
        "diastolic": 80
      },
      "origin": {
        "manualMeasurement": {
          "enteredBy": "clinician"
        }
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/1610940125/measurements/1",
        "patient": "http://localhost:8081/clinician/api/patients/1610940125"
      }
    },
    {
      "timestamp": "%s",
      "type": "weight",
      "measurement": {
        "unit": "kg",
        "value": 75
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/1610940125/measurements/2",
        "patient": "http://localhost:8081/clinician/api/patients/1610940125"
      }
    },
    {
      "timestamp": "%s",
      "type": "blood_pressure",
      "measurement": {
        "unit": "mmHg",
        "systolic": 130,
        "diastolic": 85
      },
      "origin": {
        "manualMeasurement": {
          "enteredBy": "clinician"
        }
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/1805039414/measurements/3",
        "patient": "http://localhost:8081/clinician/api/patients/1805039414"
      }
    },
    {
      "timestamp": "%s",
      "type": "weight",
      "measurement": {
        "unit": "kg",
        "value": 65
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/1805039414/measurements/4",
        "patient": "http://localhost:8081/clinician/api/patients/1805039414"
      }
    },
    {
      "timestamp": "%s",
      "type": "blood_pressure",
      "measurement": {
        "unit": "mmHg",
        "systolic": 110,
        "diastolic": 70
      },
      "origin": {
        "manualMeasurement": {
          "enteredBy": "clinician"
        }
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/2303079354/measurements/5",
        "patient": "http://localhost:8081/clinician/api/patients/2303079354"
      }
    },
    {
      "timestamp": "%s",
      "type": "weight",
      "measurement": {
        "unit": "kg",
        "value": 55
      },
      "links": {
        "measurement": "http://localhost:8081/clinician/api/patients/2303079354/measurements/6",
        "patient": "http://localhost:8081/clinician/api/patients/2303079354"
      }
    }
  ]
}
`
