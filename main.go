package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port := getPort()
	http.HandleFunc("/", printPage)

	fmt.Printf("Try to start server on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func printPage(w http.ResponseWriter, r *http.Request) {
	file := ""
	printlog(r)
	if r.Method == "GET" {
		if isMain(r) {
			file = "res/index.html"
		} else if isWSDL(r) {
			file = "xml/vvd.wsdl"
		} else {
			file = "res/error.html"
		}
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(w, f)
}

func getPort() string {
	port, exists := os.LookupEnv("VVD_PORT")
	if exists {
		return ":" + port
	}
	return ":3000"
}

func isMain(r *http.Request) bool {
	return r.URL.Path == "/"
}

func isWSDL(r *http.Request) bool {
	if r.URL.Path == "/houseCall" {
		q := r.URL.Query()
		_, ok := q["wsdl"]
		return ok
	}
	return false
}

func printlog(r *http.Request) {
	t := time.Now().Format(time.ANSIC)
	m := r.Method
	h := r.Host
	p := r.URL.Path
	fmt.Fprintf(os.Stdout, "[%s]\t%s\t%s\t%s\n", t, m, h, p)
}
