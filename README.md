pushnotification
===
[![GoDoc](https://godoc.org/github.com/Iwark/spreadsheet?status.svg)](https://godoc.org/github.com/Iwark/spreadsheet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Package `pushnotification` sends push notification via Amazon SNS

## Example

```go
package main

import (
	"log"
	"os"

	"github.com/Iwark/pushnotification"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	push := pushnotification.Service{
		AWSAccessKey:         os.Getenv("AWSAccessKey"),
		AWSAccessSecret:      os.Getenv("AWSAccessSecret"),
		AWSSNSApplicationARN: os.Getenv("AWSSNSApplicationARN"),
		AWSRegion:            os.Getenv("AWSRegion"),
	}

	err := push.Send(os.Getenv("DeviceToken"), &pushnotification.Data{
		Alert: aws.String("test message"),
		Sound: aws.String("default"),
		Badge: aws.Int(1),
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

## License

pushnotification is released under the MIT License.
