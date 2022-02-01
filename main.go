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
	http.HandleFunc("/", infoPage)
	http.HandleFunc("/houseCall", soapService)

	fmt.Printf("Try to start server on port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	printlog(r)
	sendFile("res/index.html", w)
}
func soapService(w http.ResponseWriter, r *http.Request) {
	printlog(r)
	if r.Method == "GET" && isWSDL(r) {
		sendFile("xml/vvd.wsdl", w)
	} else if r.Method == "POST" {
		GetValidatePatientInfo(w, r)
		GetHouseCallScheduleInfo(w, r)
		CreateHouseCall(w, r)
		CancelHouseCall(w, r)
	} else {
		sendFile("res/error.html", w)
	}
}

func getPort() string {
	port, exists := os.LookupEnv("VVD_PORT")
	if exists {
		return ":" + port
	}
	return ":3000"
}

func isWSDL(r *http.Request) bool {
	if r.URL.Path == "/houseCall" {
		q := r.URL.Query()
		_, ok := q["wsdl"]
		return ok
	}
	return false
}

func sendFile(file string, w http.ResponseWriter) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(w, f)
}

func sendData(data string, w http.ResponseWriter) {
	fmt.Fprintf(w, "%q", data)
}

func printlog(r *http.Request) {
	t := time.Now().Format(time.ANSIC)
	m := r.Method
	h := r.Host
	p := r.URL.Path
	fmt.Fprintf(os.Stdout, "[%s]\t%s\t%s\t%s\n", t, m, h, p)
}
