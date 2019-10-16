package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return defaultPort
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Echoing back request made to " + r.URL.Path + " to client (" + r.RemoteAddr + ")")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// allow pre-flight headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	r.Write(w)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "bar")
}

func main() {
	log.Println("starting server, listening on port 0.0.0.0:" + getServerPort())
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/foo", fooHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+getServerPort(), nil))

}
