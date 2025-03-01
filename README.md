# Beam Toolkit Go SDK

## Usage

Start using the API with your API KEY

For more examples, please refer to the `service/examples` directory

### Scraping API
```go
package main

import (
	"fmt"
	scrapeless "github.com/scrapeless-ai/scrapeless-sdk-go/service/api"
)

func main() {
	client := scrapeless.NewClient(
		scrapeless.WithAPIKey("your-api-key"),
	)

	scraper, err := client.CreateScraperTask(&scrapeless.ServiceConfig{
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

	fmt.Printf("Scraper: %+v\n", string(scraper.Res.Body()))
}
```

### Web Unlocker
```go
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

```

### Captcha Solver
```go
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
```
