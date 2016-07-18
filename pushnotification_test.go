package pushnotification

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	assert := assert.New(t)

	assert.NoError(godotenv.Load())

	svc := Service{
		AWSAccessKey:         os.Getenv("AWSAccessKey"),
		AWSAccessSecret:      os.Getenv("AWSAccessSecret"),
		AWSSNSApplicationARN: os.Getenv("AWSSNSApplicationARN"),
		AWSRegion:            os.Getenv("AWSRegion"),
	}

	err := svc.Send(os.Getenv("DeviceToken"), &Data{
		Alert: aws.String("test message"),
		Sound: aws.String("default"),
		Badge: aws.Int(1),
	})
	assert.NoError(err)
}
