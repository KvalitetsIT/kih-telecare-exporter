package measurement

/// Implementation of Measurent Interface for the clinician api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type clinicianApi struct {
	key, secret, apiUrl string
	batchSize           int
}

func (m clinicianApi) String() string {
	return fmt.Sprintf("URL: %s - key: %s - batches: %d", m.apiUrl, m.key, m.batchSize)
}

// Fetchpatient information
func (m clinicianApi) FetchPatient(person string) (PatientResult, error) {
	var patient PatientResult
	log.Debug("requesting: ", person)
	req, err := http.NewRequest(http.MethodGet, person, nil)
	if err != nil {
		log.Errorf("Error creating patient information request %v", err)
		return patient, err
	}
	addAuthorizationHeader(req)

	resp, err := client.Do(req)
	var body []byte
	if err != nil {
		log.Errorf("Error querying API  - %v", err)
	} else {
		body, _ = ioutil.ReadAll(resp.Body)
	}

	if err := json.Unmarshal(body, &patient); err != nil {
		log.Debug("Clinician said: ", string(body))
		log.Errorf("Error unmarshalling patient response  - %v", err)
		return patient, errors.Wrap(err, "Error unmarshalling patient response")
	}
	log.Debug(fmt.Sprintf("Retrieved - %+v", patient))
	return patient, nil
}

func addAuthorizationHeader(r *http.Request) {
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
}
func (m clinicianApi) FetchMeasurement(measurement string) (Measurement, error) {
	var result Measurement
	req, err := http.NewRequest(http.MethodGet, measurement, nil)
	if err != nil {
		log.Errorf("Error creating measurment request %v", err)
		return result, err
	}
	addAuthorizationHeader(req)
	resp, err := client.Do(req)
	var body []byte
	if err != nil {
		log.Errorf("Error querying API  - %v", err)
	} else {
		body, _ = ioutil.ReadAll(resp.Body)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Errorf("Error unmarshalling patient response  - %v", err)
		return result, errors.Wrap(err, "Error unmarshalling patient response")
	}
	log.Debug(fmt.Sprintf("Retrieved - %+v", result))

	return result, nil
}
func fetchMeasurements(location string, since time.Time, offset int, max int) (MeasurementResponse, error) {

	log.Debug("Since: ", since, " Offset: ", offset, " Batches: ", max)
	var result MeasurementResponse

	v := url.Values{}
	v.Set("from", since.Format(time.RFC3339))

	requestUrl := fmt.Sprintf("%s/measurements?%s&ignored=false&offset=%d&max=%d", location, v.Encode(), offset, max)

	log.Debug(requestUrl)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return result, errors.Wrap(err, "Error creating request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))

	resp, err := client.Do(req)
	var body []byte
	if err != nil {
		return result, errors.Wrap(err, "Error querying API")
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return result, fmt.Errorf("Error accessing API - server responded: %s", resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.Wrap(err, "Decoding body")
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, errors.Wrap(err, "Error converting body to type")
	}

	log.Debug("Total: ", result.Total, " Max ", result.Max, "Next ", result.Links.Next)

	return result, nil
}

func (m clinicianApi) FetchMeasurements(since time.Time, offset int) (MeasurementResponse, error) {
	// substract 2 hours
	ts := since.Add(-2 * time.Hour)

	result, err := fetchMeasurements(m.apiUrl, ts, offset, m.batchSize)
	if err != nil {
		return result, err
	}
	return result, err
}

func (m clinicianApi) CheckHealth() error {
	requestUrl := fmt.Sprintf("%s/health", m.apiUrl)
	log.Debugf("Performing health check against %s", requestUrl)

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error creating requet for %s", requestUrl))
	}
	defer resp.Body.Close()

	logrus.Infof("Got reply %s - %d, should be %d", resp.Status, resp.StatusCode, http.StatusOK)

	if http.StatusOK != resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		logrus.Infof("returning error %s", string(body))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Error reading response from %s", requestUrl))
		}

		return fmt.Errorf("Error checking status on %s - response - %s", requestUrl, string(body))

	}

	return nil
}