package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloServer responds to requests with the given URL path.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Could not get Hostname : error %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "Hello, you requested: '%s' and were served by '%s'\n", r.URL.Path, hostname)
	log.Printf("Received request for path: '%s' and were served by '%s'\n", r.URL.Path, hostname)
}

func main() {
	var addr string = ":8180"
	
	handler := http.HandlerFunc(HelloServer)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}