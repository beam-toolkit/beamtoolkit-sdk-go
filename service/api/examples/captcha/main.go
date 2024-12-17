package main

import (
	"context"
	"fmt"
	scrapeless "github.com/scrapeless-ai/scrapeless-sdk-go/service/api"
	"time"
)

func main() {
	client := scrapeless.NewClient(
		scrapeless.WithAPIKey("your-api-key"),
	)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	captcha, err := client.SolverCaptcha(timeout, &scrapeless.ServiceConfig{
		Actor: "captcha.recaptcha",
		Input: map[string]any{
			"version":    "v2",
			"pageURL":    "https://www.google.com",
			"siteKey":    "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
			"pageAction": "",
		},
	})

	if err != nil {
		println("error: ", err.Error())
		return
	}

	fmt.Printf("Captcha:%+v\n", string(captcha.Res.Body()))
}
