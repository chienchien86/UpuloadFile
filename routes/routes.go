package routes

import (

	"net/http"
	
	"UpuloadFile/api"
	
)


var SetupServer = func(appPort string) {
	
	http.Handle("/file", &api.FileApi{})
	http.ListenAndServe(":8080", nil)
}