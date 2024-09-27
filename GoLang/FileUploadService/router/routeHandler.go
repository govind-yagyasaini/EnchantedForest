package router

import (
	"github.com/gorilla/mux"
	"github.com/govindyagyasaini/fileuploadservice/download"
	"github.com/govindyagyasaini/fileuploadservice/upload"
)

// Router returns the application's router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload.UploadHandler).Methods("POST")
	r.HandleFunc("/file/{filename}", download.DownloadHandler).Methods("GET")
	return r
}
