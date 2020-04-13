package api

import (
	"net/http"
	"fmt"
	"log"

	"UpuloadFile/controller"
)

type FileApi struct {}

func (f *FileApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		controller.FileList(w, r)
	case http.MethodPost:
		controller.UploadNewFile(w, r)
	case http.MethodPut:
		controller.UpdateFileName(w, r)
	case http.MethodDelete:
		controller.DeleteFile(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to '%v' \n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to '%v' \n", r.Method, r.URL)
	}


}