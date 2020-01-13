package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/themoah/go-web-tools/core"

	"github.com/gorilla/mux"
)

// HealthCheckHanlder returns 200 and body "ok"
func HealthCheckHanlder(w http.ResponseWriter, r *http.Request) {
	log.Println("alive :-D")
	fmt.Fprintf(w, "ok")
}

// IndexHandler returns hello world
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

// EchoHandler blah-blah
func EchoHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Echoing back request made to " + r.URL.Path + " by User-Agent: " + r.UserAgent() + ")")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// allow pre-flight headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	r.Write(w)
}

// FooHandler returns "bar"
func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, core.Foo())
}

// RandomHandler returns random float
func RandomHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Echoing back request made to " + r.URL.Path + " to client (" + r.RemoteAddr + ")")
	fmt.Fprintf(w, core.Random())
}

// SecureHandler checks if connection over https
func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet")
}

// WhoisHandler checks given ipv4 against global whois db
func WhoisHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	ip := requestParams["ip"]
	// log.Println("requested ip: " + ip)
	fmt.Fprintf(w, core.WhoIS(ip))
}
