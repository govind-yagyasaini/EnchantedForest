package awssession

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoDBClient *dynamodb.DynamoDB

// InitDynamoDBClient initializes the DynamoDB client
func InitDynamoDBClient() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %s", err)
	}
	DynamoDBClient = dynamodb.New(sess)
}
