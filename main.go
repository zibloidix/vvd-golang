package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port := getPort()
	http.HandleFunc("/", handler)

	fmt.Printf("Try to start server on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	printlog(r)
	fmt.Fprintf(w, "Hello %q", r.URL.Path)
}

func getPort() string {
	port, exists := os.LookupEnv("VVD_PORT")
	if exists {
		return ":" + port
	}
	return ":3000"
}

func printlog(r *http.Request) {
	t := time.Now().Format(time.ANSIC)
	m := r.Method
	h := r.Host
	p := r.URL.Path
	fmt.Fprintf(os.Stdout, "[%s]\t%s\t%s\t%s\n", t, m, h, p)
}
