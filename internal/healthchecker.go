package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Checks URL, if
func PerformHealthCheck(client http.Client, method string, returnCode int, healthCheckUrl string) error {
	req, err := http.NewRequest(method, healthCheckUrl, nil)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error creating requet for %s", healthCheckUrl))
	}
	defer resp.Body.Close()

	logrus.Infof("Got reply %s - %d, should be %d", resp.Status, resp.StatusCode, returnCode)

	if returnCode != resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		logrus.Infof("returning error %s", string(body))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Error reading response from %s", healthCheckUrl))
		}

		return fmt.Errorf("Error checking status on %s - response - %s", healthCheckUrl, string(body))

	}
	return nil
}
