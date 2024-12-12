package main

import (
	"fmt"
	scrapeless "github.com/scrapeless-ai/scrapeless-sdk-go/service/api"
)

func main() {
	client := scrapeless.NewClient(
		scrapeless.WithAPIKey("your-api-key"),
	)

	captcha, err := client.CreateCaptchaTask(&scrapeless.ServiceConfig{
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
