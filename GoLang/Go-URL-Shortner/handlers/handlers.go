package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/govind-yagyasaini/EnchantedForest/Go-URL-Shortner/awssession"
	"github.com/govind-yagyasaini/EnchantedForest/Go-URL-Shortner/config"
	"github.com/govind-yagyasaini/EnchantedForest/Go-URL-Shortner/models"
	"golang.org/x/exp/rand"
)

// HandleShorten handles both GET (form display) and POST (URL shortening)
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		formHTML := `
		<h2>URL Shortener</h2>
		<form method="post" action="/shorten">
			<input type="text" name="url" placeholder="Enter a URL to shorten">
			<input type="submit" value="Shorten">
		</form>
		`
		fmt.Fprint(w, formHTML)
		return
	}

	if r.Method == http.MethodPost {
		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "Please provide the URL to shorten.", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
			originalURL = "http://" + originalURL
		}

		_, err := url.ParseRequestURI(originalURL)
		if err != nil {
			http.Error(w, "Input URL is invalid.", http.StatusBadRequest)
			return
		}

		input := &dynamodb.QueryInput{
			TableName:              aws.String(config.DynamoDBTableName),
			IndexName:              aws.String("originalurl-index"),
			KeyConditionExpression: aws.String("originalurl = :url"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":url": {
					S: aws.String(originalURL),
				},
			},
		}

		queryResult, err := awssession.DynamoDBClient.Query(input)
		if err != nil {
			log.Printf("Got error querying DynamoDB: %s", err)
			http.Error(w, "Error checking existing URL", http.StatusInternalServerError)
			return
		}

		if len(queryResult.Items) > 0 {
			existingItem := models.Item{}
			err = dynamodbattribute.UnmarshalMap(queryResult.Items[0], &existingItem)
			if err != nil {
				log.Printf("Failed to unmarshal existing item: %s", err)
				http.Error(w, "Error processing existing URL", http.StatusInternalServerError)
				return
			}

			shortenedURL := fmt.Sprintf("http://localhost:4003/short/%s", existingItem.Shortkey)
			w.Header().Set("Content-Type", "text/html")
			responseHTML := fmt.Sprintf(`
			<h2>URL Shortener</h2>
			<p>Original URL: %s</p>
			<p>Shortened URL: <a href="%s">%s</a></p>
			<form method="post" action="/shorten">
				<input type="text" name="url" placeholder="Enter a URL">
				<input type="submit" value="Shorten">
			</form>
			`, originalURL, shortenedURL, shortenedURL)

			fmt.Fprint(w, responseHTML)
			return
		}

		shortkey_ := GenerateShortKeys()
		item := models.Item{Shortkey: shortkey_, Originalurl: originalURL}
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			log.Fatalf("Got error marshalling the url: %s", err)
		}

		putInput := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(config.DynamoDBTableName),
		}

		_, err = awssession.DynamoDBClient.PutItem(putInput)
		if err != nil {
			log.Fatalf("Got error calling PutItem: %s", err)
		}

		shortenedURL := fmt.Sprintf("http://localhost:4003/short/%s", shortkey_)
		w.Header().Set("Content-Type", "text/html")
		responseHTML := fmt.Sprintf(`
		<h2>URL Shortener</h2>
		<p>Original URL: %s</p>
		<p>Shortened URL: <a href="%s">%s</a></p>
		<form method="post" action="/shorten">
			<input type="text" name="url" placeholder="Enter a URL">
			<input type="submit" value="Shorten">
		</form>
		`, originalURL, shortenedURL, shortenedURL)

		fmt.Fprint(w, responseHTML)
	}
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	result, err := awssession.DynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.DynamoDBTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"shortkey": {
				S: aws.String(shortKey),
			},
		},
	})

	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
		http.Error(w, "Error fetching the original URL", http.StatusInternalServerError)
		return
	}

	if result.Item == nil {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	item := models.Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Printf("Failed to unmarshal DynamoDB result: %s", err)
		http.Error(w, "Error processing the result", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, item.Originalurl, http.StatusMovedPermanently)
}

func GenerateShortKeys() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	rand.Seed(uint64(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
