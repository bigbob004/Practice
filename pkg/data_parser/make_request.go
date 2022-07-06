package data_parser

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func MakeRequest(siteURL string, headers map[string]string, timeout int, attempts int) ([]byte, error) {
	body := io.Reader(nil)
	req, err := http.NewRequest(http.MethodGet, siteURL, body)
	if err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	// Use the default timeout if the timeout parameter isn't configured.
	reqTimeout := 10 * time.Second
	if timeout != 0 {
		reqTimeout = time.Duration(timeout) * time.Second
	}
	//TODO: избавиться от кастомного http-клиента Use default http Client.
	httpClient := &http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       reqTimeout,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 502 {
			i := 0
			for resp.StatusCode == 502 && i <= attempts {
				time.Sleep(5 * time.Millisecond)
				resp, err = httpClient.Do(req)
				if err != nil {
					return nil, err
				}
				attempts++
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Http status code is %d", resp.StatusCode))
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
