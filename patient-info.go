package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"text/template"
)

type PatientInfoResponse struct {
	SessionID    string
	PatientID    int
	MOID         int
	MOOID        string
	MOName       string
	MOAddress    string
	MOPhone      string
	ResourceID   string
	ResourceName string
}

func GetValidatePatientInfo(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "GetValidatePatientInfo" {
		templateFile := respBase + action + "/" + respOk
		envelope, err := ParseSoapEnvelope(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr02
		}
		patientInfo, err := getSQLData(envelope)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr01
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr02
		}
		setSessionID(&patientInfo, envelope)
		data, err := os.ReadFile(templateFile)
		if err != nil {
			panic(err)
		}
		t := template.New("tmp")
		t.Parse(string(data))
		w.Header().Add("Content-Type", "application/xml")
		t.Execute(w, patientInfo)
	}
}

func getSQLData(e Envelope) (PatientInfoResponse, error) {
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	p := PatientInfoResponse{}
	if err != nil {
		return p, err
	}
	defer db.Close()
	sqlFile, err := os.ReadFile("./sql/GetValidatePatientInfo.sql")
	if err != nil {
		return p, err
	}
	city, street, house := getParamsFromRequest(e)
	row := db.QueryRow(string(sqlFile), city, street, house)
	err = row.Scan(&p.PatientID, &p.MOID, &p.MOOID, &p.MOName, &p.MOAddress, &p.MOPhone, &p.ResourceID, &p.ResourceName)
	if err != nil {
		return p, err
	}
	return p, nil
}

func getParamsFromRequest(e Envelope) (string, string, int) {
	city := e.Body.GetValidatePatientInfoRequest.Adr_City
	street := e.Body.GetValidatePatientInfoRequest.Adr_Street
	house := e.Body.GetValidatePatientInfoRequest.Adr_House
	return city, street, house
}

func setSessionID(p *PatientInfoResponse, e Envelope) {
	p.SessionID = e.Body.GetValidatePatientInfoRequest.Session_ID
}
