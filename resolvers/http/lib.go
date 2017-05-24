package http

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gregjones/httpcache"
)

var baseURL = "https://esi.tech.ccp.is/latest/"
var dataSource = "tranquility"

var client = httpcache.NewMemoryCacheTransport().Client()

func doRequest(method string, auth *string, url string, queryParams *map[string]string, params ...interface{}) (*http.Response, error) {
	formattedTarget := fmt.Sprintf(url, params...)

	var paramMap map[string]string
	if queryParams != nil {
		paramMap = *queryParams
	} else {
		paramMap = make(map[string]string)
	}

	paramMap["datasource"] = dataSource
	paramMap["language"] = "en-us"
	queryString := buildQueryString(paramMap)

	fullURL := fmt.Sprintf("%s%s?%s", baseURL, formattedTarget, queryString)
	req, _ := http.NewRequest(method, fullURL, nil)

	if auth != nil {
		req.Header.Set("Authorization", *auth)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got unexpected StatusCode %d back from ESI", resp.StatusCode)
	}

	return resp, nil
}

func buildQueryString(params map[string]string) string {
	vals := make(url.Values)
	for key, value := range params {
		vals.Add(key, value)
	}

	return vals.Encode()
}

// MakeQuery makes requests to ESI endpoints using provided query string parameters
func MakeQuery(url string, queryParams map[string]string, urlParams ...interface{}) (*http.Response, error) {
	return doRequest("GET", nil, url, &queryParams, urlParams...)
}

// MakeRequest requests ESI endpoints using a caching http client
func MakeRequest(url string, params ...interface{}) (*http.Response, error) {
	return doRequest("GET", nil, url, nil, params...)
}

// MakeAuthorizedRequest makes requests for ESI endpoints that require authorization
func MakeAuthorizedRequest(auth string, url string, params ...interface{}) (*http.Response, error) {
	return doRequest("GET", &auth, url, nil, params...)
}

// MakeAuthorizedQuery makes requests to ESI endpoints that require authoirzation using
// the provided query string parameters
func MakeAuthorizedQuery(auth string, queryParams map[string]string, url string, urlParams ...interface{}) (*http.Response, error) {
	return doRequest("GET", &auth, url, &queryParams, urlParams...)
}
