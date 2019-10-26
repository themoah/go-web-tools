package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/likexian/whois-go"
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

// FooHandler returns "bar"
func FooHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "bar")
}

// RandomHandler returns random float
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Echoing back request made to " + r.URL.Path + " to client (" + r.RemoteAddr + ")")
	i := rand.Float64()
	iStr := fmt.Sprintf("%f", i)
	fmt.Fprintf(w, iStr)
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

func ipWhoisHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	ip := requestParams["ip"]
	log.Println("requested ip: " + ip)
	result, err := whois.Whois(ip)
	if err == nil {
		fmt.Fprintf(w, result)
	}
}

func main() {
	log.Println("starting server, listening on port 0.0.0.0:" + getServerPort())

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/echo", echoHandler).Methods("GET")
	r.HandleFunc("/foo", FooHandler).Methods("GET")
	r.HandleFunc("/random", RandomHandler).Methods("GET")
	r.HandleFunc("/secure", secureHandler).Methods("GET").Schemes("https")
	r.HandleFunc("/ip/{ip}", ipWhoisHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+getServerPort(), r))

}
