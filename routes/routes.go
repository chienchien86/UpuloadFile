package routes

import (

	"log"
	"net/http"
	"time"

	"UpuloadFile/controller"
	
)


type Middleware func(http.Handler) http.Handler

func Chain(f http.Handler, mmap ...Middleware) http.Handler {
	for _, m := range mmap {
	   f = m(f)
	}
	return f
}

func Method(m string) Middleware {
	return func(f http.Handler) http.Handler {
	   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		  log.Println(r.URL.Path)
		  if r.Method != m {
			 http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			 return
		  }
		  f.ServeHTTP(w, r)
	   })
	}
 
}

func Log() Middleware {
	return func(f http.Handler) http.Handler {
	   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		  //log.Println(r.URL.Path)
		  // Do middleware things
		  start := time.Now()
		  defer func() { log.Println(r.URL.Path, time.Since(start)) }()
		  f.ServeHTTP(w, r)
	   })
	}
}


var SetupServer = func(appPort string) {
	http.Handle("/filelist", Chain(http.HandlerFunc(controller.FileList), Method("GET"), Log()))
	http.Handle("/filesremove", Chain(http.HandlerFunc(controller.DeleteFile), Method("DELETE"), Log()))
	http.Handle("/fileupload", Chain(http.HandlerFunc(controller.UploadNewFile), Method("POST"), Log()))
	http.Handle("/filerename", Chain(http.HandlerFunc(controller.UpdateFileName), Method("PUT"), Log()))
	http.ListenAndServe(":8080", nil)
}