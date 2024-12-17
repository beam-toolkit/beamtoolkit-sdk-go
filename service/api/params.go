package api

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

var (
	NotConfiguredError   = SdkError{errors.New("scrapeless client is not configured")}
	NotFoundTaskIdError  = SdkError{errors.New("task id not found")}
	SolverCaptchaTimeout = SdkError{errors.New("solver captcha timeout")}
)

type ServiceConfig struct {
	Actor string         `json:"actor"`
	Input map[string]any `json:"input"`
	Proxy map[string]any `json:"proxy"`
}

type CaptchaResult struct {
	State   string `json:"state"`
	Success bool   `json:"success"`
	TaskId  string `json:"taskId"`
}

type CaptchaTaskResult struct {
	Actor      string `json:"actor"`
	CreateTime int64  `json:"createTime"`
	Elapsed    int    `json:"elapsed"`
	Solution   struct {
		Token string `json:"token"`
	} `json:"solution"`
	Success bool   `json:"success"`
	TaskId  string `json:"taskId"`
}

type Response struct {
	Res *resty.Response
}

type SdkError struct {
	error
}
