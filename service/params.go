package scrapeless

import (
	"errors"
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
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type SdkError struct {
	error
}
