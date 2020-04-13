package controller


import (
	"net/http"
    "log"
	"os"
	"fmt"
	"io/ioutil"
	"time"
	"io"
	
	"UpuloadFile/RequestLimit"
	u "UpuloadFile/utils"

)

var requestLimit = RequestLimit.NewRequestLimitService(10 * time.Second, 10)

var UploadNewFile = func(w http.ResponseWriter, r *http.Request) {
	if requestLimit.IsAvailable() {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("filename")
		fileBytes, err := ioutil.ReadAll(file)
    	if err != nil {
			fmt.Println(err)
			return 
    	}
    	err = ioutil.WriteFile("./static/" + handler.Filename, fileBytes, 0644)
    	if err != nil {
        	panic(err)
    	}
		resp := u.Message(true, "Successful")
		resp["uploadfilename"] = handler.Filename
		u.Respond(w, resp)
	}else{
		fmt.Println("Reach request limiting!")
		io.WriteString(w, "Reach request limit!\n")
	}

}



var DeleteFile = func(w http.ResponseWriter, r *http.Request) {
	if requestLimit.IsAvailable() {
		name := r.FormValue("filename")
		err := os.Remove("./static/" + name)
    	if err != nil {
        	fmt.Println(err)
        	return
		}
		resp := u.Message(true, "delete Successful")
		u.Respond(w, resp)
	}else{
		fmt.Println("Reach request limiting!")
		io.WriteString(w, "Reach request limit!\n")
	}
}

var FileList = func(w http.ResponseWriter, r *http.Request) {
	if requestLimit.IsAvailable() {
		file, err := os.Open("./static")
		if err != nil {
        	log.Fatalf("failed opening directory: %s", err)
    	}
		list,_ := file.Readdirnames(0) 
		resp := u.Message(true, "Successful")
		resp["data"] = list
		u.Respond(w, resp)
	}else{
		fmt.Println("Reach request limiting!")
		io.WriteString(w, "Reach request limit!\n")
	}
}

var UpdateFileName = func(w http.ResponseWriter, r *http.Request) {
	if requestLimit.IsAvailable() {
		old :=  "./staic/" + r.FormValue("oldfilename")
		new :=  "./staic/" + r.FormValue("newfilename")
			
		err := os.Rename( old, new)
		if err != nil {
        	fmt.Println(err)
        	return 
		}
		resp := u.Message(true, "rename file Successful")
		u.Respond(w, resp)
		
	} else{
		fmt.Println("Reach request limiting!")
		io.WriteString(w, "Reach request limit!\n")
	}
	
}