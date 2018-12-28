package http

import (
	"net/http"
	"net/url"
	"time"
)

var client = &http.Client{}

// SetTimeout sets timeout for HTTP requests in seconds.
func SetTimeout(timeout int) {
	client.Timeout = time.Duration(int(time.Second) * timeout)
}

// Get emits a HTTP GET request for a given URL. Captures the status code, status and outcome of the call.
// Returns with information about the response.
func Get(link string) HttpResponse {
	result := HttpResponse{Url: link}
	url, err := url.ParseRequestURI(link)

	if err != nil {
		result.Error = err
		return result
	}

	resp, err := client.Get(url.String())

	if err != nil {
		result.Error = err
		return result
	}

	result.StatusCode = resp.StatusCode
	result.Status = resp.Status

	if resp.StatusCode == 200 {
		result.Success = true
	}

	resp.Body.Close()
	return result
}

// HttpResponse represents HTTP response information.
type HttpResponse struct {
	Url        string
	Success    bool
	StatusCode int
	Status     string
	Error      error
}
