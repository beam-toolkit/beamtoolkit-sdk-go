package api

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

var (
	NotConfiguredError  = SdkError{errors.New("scrapeless client is not configured")}
	NotFoundTaskIdError = SdkError{errors.New("task id not found")}
)

type ServiceConfig struct {
	Actor string         `json:"actor"`
	Input map[string]any `json:"input"`
	Proxy map[string]any `json:"proxy"`
}

type Response struct {
	Res *resty.Response
}

type SdkError struct {
	error
}
