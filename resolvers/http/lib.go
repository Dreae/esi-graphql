package http

import (
	"fmt"
	"net/http"
)

var baseURL = "https://esi.tech.ccp.is/latest/"
var dataSource = "tranquility"

func doRequest(method string, auth *string, url string, params ...interface{}) (*http.Response, error) {
	formattedTarget := fmt.Sprintf(url, params...)
	fullURL := fmt.Sprintf("%s%s?datasource=%s", baseURL, formattedTarget, dataSource)
	req, _ := http.NewRequest(method, fullURL, nil)

	if auth != nil {
		req.Header.Set("Authorization", *auth)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got unexpected StatusCode %d back from ESI", resp.StatusCode)
	}

	return resp, nil
}

func MakeRequest(url string, params ...interface{}) (*http.Response, error) {
	return doRequest("GET", nil, url, params...)
}

func MakeAuthorizedRequest(auth string, url string, params ...interface{}) (*http.Response, error) {
	return doRequest("GET", &auth, url, params...)
}
