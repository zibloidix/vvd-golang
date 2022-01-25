package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := getPort()
	http.HandleFunc("/", handler)

	fmt.Printf("Try to start server on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", r.URL.Path)
}

func getPort() string {
	port, exists := os.LookupEnv("VVD_PORT")
	if exists {
		return ":" + port
	}
	return ":3000"
}
