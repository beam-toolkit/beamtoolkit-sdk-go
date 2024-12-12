package main

import (
	"fmt"
	scrapeless "github.com/scrapeless-ai/scrapeless-sdk-go/service/api"
)

func main() {
	client := scrapeless.NewClient(
		scrapeless.WithAPIKey("your-api-key"),
	)

	unlocker, err := client.Unlocker(&scrapeless.ServiceConfig{
		Actor: "unlocker.webunlocker",
		Input: map[string]any{
			"url":           "https://www.scrapeless.com",
			"proxy_country": "ANY",
			"method":        "GET",
			"redirect":      false,
		},
	})

	if err != nil {
		println("error: ", err.Error())
		return
	}

	fmt.Printf("Scraper: %+v\n", string(unlocker.Res.Body()))
}
