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

	r.ParseMultipartForm(10 << 20) // Limit to 10MB

	// Retrieve the file from form-data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		log.Printf("Error retrieving the file: %v", err)
		return
	}
	defer file.Close()

	var session = awssession.GetAWSConnection()
	uploader := s3manager.NewUploader(session)
	_, err = uploader.Upload(&s3manager.UploadInput{Bucket: aws.String(config.AWS_S3_BUCKET), Key: aws.String(handler.Filename), Body: file})
	if err != nil {
		log.Fatal(err)
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}

//For uploading file to local system

// func UploadHandler(w http.ResponseWriter, r *http.Request) {

// 	r.ParseMultipartForm(10 << 20) // Limit to 10MB

// 	// Retrieve the file from form-data
// 	file, handler, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
// 		log.Printf("Error retrieving the file: %v", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create an uploads directory if it doesn't exist
// 	uploadDir := "./uploads"
// 	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
// 		err = os.Mkdir(uploadDir, 0755)
// 		if err != nil {
// 			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
// 			log.Printf("Error creating upload directory: %v", err)
// 			return
// 		}
// 	}

// 	//Define the file path
// 	filePath := uploadDir + "/" + handler.Filename
// 	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		http.Error(w, "Error saving the file", http.StatusInternalServerError)
// 		log.Printf("Error saving the file: %v", err)
// 		return
// 	}
// 	defer f.Close()

// 	//Copy the uploaded file to the server
// 	_, err = io.Copy(f, file)
// 	if err != nil {
// 		http.Error(w, "Error writing the file", http.StatusInternalServerError)
// 		log.Printf("Error writing the file: %v", err)
// 		return
// 	}

// 	// Send a success response
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
// }
