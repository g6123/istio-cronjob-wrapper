package pkg

import (
	"fmt"
	"net/http"
	"time"
)

func WaitEnvoyReady(url string, connect_timeout int, max_retry int) error {
	client := &http.Client{
		Timeout:       time.Duration(connect_timeout) * time.Second,
		CheckRedirect: checkRedirect,
	}

	retry := 0

	for {
		resp, err := client.Get(url)
		if err == nil && resp.StatusCode != 200 {
			err = fmt.Errorf("GET \"%s\": status: %s", url, resp.Status)
		}

		if err == nil {
			return nil
		}

		if retry++; retry > max_retry {
			return err
		}

		Logger.Debug(err)
		Logger.Warningf("envoy not ready. retry after 1s. (%d/%d)", retry, max_retry)
		time.Sleep(1 * time.Second)
	}
}

func KillEnvoy(url string) error {
	req, err := http.NewRequest("POST", url, http.NoBody)
	if err != nil {
		return err
	}

	client := http.Client{CheckRedirect: checkRedirect}
	_, err = client.Do(req)
	return err
}
