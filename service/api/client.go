package api

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	cfg  options
	http *resty.Client
}

const baseUrl = "https://api.scrapeless.com/api/v1"

func NewClient(opts ...Options) *Client {
	client := &Client{cfg: options{}}

	for _, opt := range opts {
		opt.apply(&client.cfg)
	}

	client.http = resty.New().
		SetBaseURL(baseUrl).
		SetCloseConnection(true).
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second)

	return client
}

func (c *Client) isConfigured() bool {
	return c.cfg.apiKey != ""
}

func (c *Client) worker(url string, params *ServiceConfig) (*Response, error) {
	if !c.isConfigured() {
		return nil, NotConfiguredError
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	workUrl := baseUrl + url
	resp, err := c.http.R().
		SetHeader("x-api-token", c.cfg.apiKey).
		SetBody(jsonData).
		Post(workUrl)

	if err != nil {
		return nil, err
	}

	return &Response{Res: resp}, nil
}

func (c *Client) get(url string, taskId string) (*Response, error) {
	if !c.isConfigured() {
		return nil, NotConfiguredError
	}
	if taskId == "" {
		return nil, NotFoundTaskIdError
	}

	workUrl := baseUrl + url + "/" + taskId
	resp, err := c.http.R().
		SetHeader("x-api-token", c.cfg.apiKey).
		Get(workUrl)

	if err != nil {
		return nil, err
	}

	return &Response{Res: resp}, nil
}

func (c *Client) CreateScraperTask(params *ServiceConfig) (*Response, error) {
	resp, err := c.worker("/scraper/requestv2", params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Unlocker(params *ServiceConfig) (*Response, error) {
	resp, err := c.worker("/unlocker/request", params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateCaptchaTask(params *ServiceConfig) (*Response, error) {
	resp, err := c.worker("/createTask", params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetCaptchaResult(taskId string) (*Response, error) {
	resp, err := c.get("/getTaskResult", taskId)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetScraperResult(taskId string) (*Response, error) {
	resp, err := c.get("/scraper/result", taskId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
