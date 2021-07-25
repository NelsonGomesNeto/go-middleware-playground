package main

import (
	"fmt"
	"go-middleware-playground/middlewares"
	"log"
	"net/http"
	"time"
)

func HelloHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("おはよう"))
}

func CurrentTimeHanlder(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("Current time: %v", currentTime)))
}

func main() {
	hostname := "localhost"
	port := 4000
	serverHost := fmt.Sprintf("%s:%d", hostname, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHanlder)
	mux.HandleFunc("/v1/time", CurrentTimeHanlder)

	wrappedMux := middlewares.NewLogger(mux)

	log.Printf("Server listening at %s", serverHost)
	log.Fatal(http.ListenAndServe(serverHost, wrappedMux))
}
