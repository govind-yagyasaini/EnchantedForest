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

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file name.", err)
	}

	var session = awssession.GetAWSConnection()

	downloader := s3manager.NewDownloader(session)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(config.AWS_S3_BUCKET),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Fatal("Error while downloading file.:", err)
	}

	http.ServeFile(w, r, filename)

}

// Download files from local system

// func DownloadHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	filename := vars["filename"]

// 	// Open the file for reading
// 	filePath := "./uploads/" + filename
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		http.Error(w, "File not found", http.StatusNotFound)
// 		return
// 	}
// 	defer file.Close()

// 	// Set headers for file download
// 	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
// 	w.Header().Set("Content-Type", "application/octet-stream")

// 	// Stream the file to the client
// 	//io.Copy(w, file)
// 	http.ServeFile(w, r, filePath)

// }
