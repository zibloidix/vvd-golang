package main

import (
	"os"
	"github.com/aymerick/raymond"
	"net/http"
)

func GetValidatePatientInfo(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "GetValidatePatientInfo" {
		data, _ := os.ReadFile("./xml/GetValidatePatientInfo/Response_Ok.xml")
		ctx := map[string]string{
			"Patient_Id": "908765",
		}
		render, _ := raymond.Render(string(data), ctx)
		sendData(render, w)
	}
}
