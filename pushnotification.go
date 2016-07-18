package pushnotification

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Service is the main entry point into using this package.
type Service struct {
	AWSAccessKey         string
	AWSAccessSecret      string
	AWSSNSApplicationARN string
	AWSRegion            string
}

// Data is the data of the sending pushnotification.
type Data struct {
	Alert *string     `json:"alert,omitempty"`
	Sound *string     `json:"sound,omitempty"`
	Data  interface{} `json:"custom_data"`
	Badge *int        `json:"badge,omitempty"`
}

// Send sends a push notification
func (s *Service) Send(deviceToken string, data *Data) (err error) {

	svc := sns.New(session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s.AWSAccessKey, s.AWSAccessSecret, ""),
		Region:      aws.String(s.AWSRegion),
	}))

	resp, err := svc.CreatePlatformEndpoint(&sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(s.AWSSNSApplicationARN),
		Token: aws.String(deviceToken),
	})
	if err != nil {
		return
	}

	m, err := newMessageJSON(data)
	if err != nil {
		return
	}

	input := &sns.PublishInput{
		Message:          aws.String(m),
		MessageStructure: aws.String("json"),
		TargetArn:        aws.String(*resp.EndpointArn),
	}
	_, err = svc.Publish(input)
	return
}

type message struct {
	APNS        string `json:"APNS"`
	APNSSandbox string `json:"APNS_SANDBOX"`
	Default     string `json:"default"`
	GCM         string `json:"GCM"`
}

type iosPush struct {
	APS Data `json:"aps"`
}

type gcmPush struct {
	Message *string     `json:"message,omitempty"`
	Custom  interface{} `json:"custom"`
	Badge   *int        `json:"badge,omitempty"`
}

type gcmPushWrapper struct {
	Data gcmPush `json:"data"`
}

func newMessageJSON(data *Data) (m string, err error) {
	b, err := json.Marshal(iosPush{
		APS: *data,
	})
	if err != nil {
		return
	}
	payload := string(b)

	b, err = json.Marshal(gcmPushWrapper{
		Data: gcmPush{
			Message: data.Alert,
			Custom:  data.Data,
			Badge:   data.Badge,
		},
	})
	if err != nil {
		return
	}
	gcm := string(b)

	pushData, err := json.Marshal(message{
		Default:     *data.Alert,
		APNS:        payload,
		APNSSandbox: payload,
		GCM:         gcm,
	})
	if err != nil {
		return
	}
	m = string(pushData)
	return
}
