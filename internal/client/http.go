package client

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// RestClient wraps resty to make HTTP requests to the ecommerce API.
type RestClient struct {
	client       *resty.Client
	baseURL      string
	defaultToken string
	useToken     bool
	logger       *logrus.Logger
}

// NewRestClient creates a new RestClient configured with the base URL and auth token.
func NewRestClient(baseURL, defaultToken string, logger *logrus.Logger) *RestClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetTimeout(30*time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	// Log method, URL, and status after every response
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		logger.WithFields(logrus.Fields{
			"method": resp.Request.Method,
			"url":    resp.Request.URL,
			"status": resp.StatusCode(),
		}).Info("HTTP response received")
		return nil
	})

	rc := &RestClient{
		client:       client,
		baseURL:      baseURL,
		defaultToken: defaultToken,
		logger:       logger,
	}

	logger.WithFields(logrus.Fields{
		"baseURL":  baseURL,
		"hasToken": rc.useToken,
	}).Info("HTTP client initialized")

	return rc
}

func (c *RestClient) PrepareRequest() *resty.Request {
	request := c.client.R()
	if c.useToken && c.defaultToken != "" {
		request.SetAuthToken(c.defaultToken)
	}
	return request
}

func (c *RestClient) WithToken() *RestClient {
	clone := *c
	clone.useToken = true
	return &clone
}

// Get sends a GET request to the given path and returns the response body.
func (c *RestClient) Get(path string, queryParams map[string]string) ([]byte, error) {
	req := c.PrepareRequest()
	if len(queryParams) > 0 {
		req.SetQueryParams(queryParams)
	}
	resp, err := req.Get(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Post sends a POST request with a JSON body and returns the response body.
func (c *RestClient) Post(path string, body interface{}) ([]byte, error) {
	resp, err := c.PrepareRequest().
		SetBody(body).
		Post(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Put sends a PUT request with a JSON body and returns the response body.
func (c *RestClient) Put(path string, body interface{}) ([]byte, error) {
	resp, err := c.PrepareRequest().
		SetBody(body).
		Put(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Delete sends a DELETE request and returns the response body.
func (c *RestClient) Delete(path string) ([]byte, error) {
	resp, err := c.PrepareRequest().
		Delete(path)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
