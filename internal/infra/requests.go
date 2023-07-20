package infra

import (
	"bytes"
	"io"
	"net/http"
)

type Requests struct{}

var httpClient = &http.Client{}

func (r Requests) Post(url string, body []byte, headers map[string]string) (response []byte, err error) {
	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	stream, _ := io.ReadAll(res.Body)

	return stream, nil
}
