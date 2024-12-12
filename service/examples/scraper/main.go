package main

import (
	"fmt"
	scrapeless "github.com/scrapeless-ai/scrapeless-sdk-go/service"
)

func main() {
	client := scrapeless.NewClient(
		scrapeless.WithAPIKey("your-api-key"),
	)

	scraper, err := client.Scraper(&scrapeless.ServiceConfig{
		Actor: "scraper.google.trends",
		Input: map[string]any{
			"keywords": "iphone14,iphone13",
			"geo":      "",
			"time":     "now 1-d",
			"category": "0",
			"property": "",
		},
	})

	if err != nil {
		println("error: ", err.Error())
		return
	}

	fmt.Printf("Scraper: %+v\n", scraper.Message)
}
