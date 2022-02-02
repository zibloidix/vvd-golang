package main

import (
	"os"
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"github.com/aymerick/raymond"
	"net/http"
)

type PatientInfoResponse struct {
	PatientID int
	MOID int
	MOOID int
	MOName string
	MOAddress string
	MOPhone string
	ResourceID string
	ResourceName string
}

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
