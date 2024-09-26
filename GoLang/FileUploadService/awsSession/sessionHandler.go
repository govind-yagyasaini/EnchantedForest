package awssession

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/govindyagyasaini/fileuploadservice/config"
)

var (
	awsSession *session.Session
	once       sync.Once
)

// GetAWSConnection returns a singleton session object
func GetAWSConnection() *session.Session {
	once.Do(func() {
		var err error
		awsSession, err = session.NewSession(&aws.Config{Region: aws.String(config.AWS_S3_REGION)})
		if err != nil {
			log.Fatal(err)
		}
	})
	return awsSession
}
