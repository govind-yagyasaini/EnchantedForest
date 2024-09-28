package models

type Item struct {
	Shortkey    string `json:"shortkey" dynamodbav:"shortkey"`
	Originalurl string `json:"originalurl" dynamodbav:"originalurl"`
}
