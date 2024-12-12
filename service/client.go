package scrapeless

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Client struct {
	cfg  options
	http *fasthttp.Client
}

const baseUrl = "https://api.scrapeless.com/api/v2"

func NewClient(opts ...Options) *Client {
	client := &Client{cfg: options{}}

	for _, opt := range opts {
		opt.apply(&client.cfg)
	}

	client.http = new(fasthttp.Client)

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
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(workUrl)
	req.Header.Set("x-api-token", c.cfg.apiKey)
	req.SetBody(jsonData)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	err = c.http.Do(req, resp)

	if err != nil {
		return nil, err
	}

	var body *Response
	bodyErr := json.Unmarshal(resp.Body(), &body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	return body, nil
}

func (c *Client) get(url string, taskId string) (*Response, error) {
	if !c.isConfigured() {
		return nil, NotConfiguredError
	}
	if taskId == "" {
		return nil, NotFoundTaskIdError
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI(baseUrl + url + "/" + taskId)
	req.Header.Set("x-api-token", c.cfg.apiKey)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	err := c.http.Do(req, resp)
	if err != nil {
		return nil, err
	}

	var body *Response
	bodyErr := json.Unmarshal(resp.Body(), &body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	return body, nil
}

func (c *Client) Scraper(params *ServiceConfig) (*Response, error) {
	resp, err := c.worker("/scraper/request", params)
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

func (c *Client) Captcha(params *ServiceConfig) (*Response, error) {
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
