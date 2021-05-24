package aws

import (
	"list_product/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetAwsSession -- initialize and loads aws session
func GetAwsSession(awsRegion string) *session.Session {
	return session.Must(session.NewSession(
		&aws.Config{
			Endpoint: aws.String(config.ENV.AwsEndpoint),
			Region:   aws.String(config.ENV.AwsRegion),
		}))
}
