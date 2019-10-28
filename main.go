package main

import (
	"log"
	"net/http"
	"os"

	"github.com/themoah/go-web-tools/routes"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
		log.Printf("Defaulting to port %s", serverPort)
	}

	log.Println("starting server, listening on port 0.0.0.0:" + serverPort)

	r := mux.NewRouter()

	r.HandleFunc("/", routes.IndexHandler)
	r.HandleFunc("/echo", routes.EchoHandler).Methods("GET")
	r.HandleFunc("/foo", routes.FooHandler).Methods("GET")
	r.HandleFunc("/random", routes.RandomHandler).Methods("GET")
	r.HandleFunc("/secure", routes.SecureHandler).Methods("GET").Schemes("https")
	r.HandleFunc("/ip/{ip}", routes.IpWhoisHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+serverPort, r))

}
