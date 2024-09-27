package upload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	awssession "github.com/govindyagyasaini/fileuploadservice/awsSession"
	"github.com/govindyagyasaini/fileuploadservice/config"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form-data, max 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to process request", http.StatusBadRequest)
		log.Println("ParseMultipartForm error:", err)
		return
	}

	// Retrieve file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found", http.StatusInternalServerError)
		log.Println("Error retrieving file:", err)
		return
	}
	defer file.Close()

	// Get AWS session and upload file
	session := awssession.GetAWSConnection()
	uploader := s3manager.NewUploader(session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWS_S3_BUCKET),
		Key:    aws.String(handler.Filename),
		Body:   file,
	})
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		log.Println("S3 upload error:", err)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}
