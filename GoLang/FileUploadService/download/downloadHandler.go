package download

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	awssession "github.com/govindyagyasaini/fileuploadservice/awsSession"
	"github.com/govindyagyasaini/fileuploadservice/config"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	// Create a file on the local system
	f, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		log.Println("File creation error:", err)
		return
	}
	defer f.Close()

	// Initiate S3 session and downloader
	session := awssession.GetAWSConnection()
	downloader := s3manager.NewDownloader(session)

	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(config.AWS_S3_BUCKET),
		Key:    aws.String(filename),
	})
	if err != nil {
		http.Error(w, "Unable to download file", http.StatusInternalServerError)
		log.Println("S3 download error:", err)
		return
	}

	// Serve file back to the client
	http.ServeFile(w, r, filename)
}
