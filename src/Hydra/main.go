package main

import (
	"Hydra/hlogger"
	"fmt"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request){
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra software systesm")

	logger.Println("recieved an http an Request on root url")
}
