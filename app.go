package main

import (
	"os"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

type Configuration struct {
	port 		string
	dbName 		string
	dbPassword 	string
}

func NewConfiguration() *Configuration {
	conf := Configuration{}
	conf.load()
	return &conf;
}


func (c *Configuration) load() {
	c.getPort()
}

func (c *Configuration) getPort() {
	pstr := os.Getenv("PORT")
	port, err := strconv.Atoi(pstr)
	if err != nil {
		port = 0
		log.Printf("failed converting '%s'\n", pstr)
	}

	if (port == 0) {
		port = 8080
		log.Printf("invalid port; fall back to default\n")
	}

	c.port = ":" + strconv.Itoa(port)
	log.Println("using port: " + c.port)
}


func main() {

	// load configuration
	config := NewConfiguration()

	// connect to db
	// TODO

	apiHandler := func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/api/", apiHandler)
	mux.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	http.ListenAndServe(config.port, mux)
}
