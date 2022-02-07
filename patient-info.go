package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"text/template"
)

type PatientInfoResponse struct {
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
		data, _ := os.ReadFile("./xml/GetValidatePatientInfo/Response_Ok.xml")
		p := getSQLData()
		e := ParseSoapEnvelope(r.Body)
		fmt.Println(e)
		t := template.New("tmp")
		t.Parse(string(data))
		w.Header().Add("Content-Type", "application/xml")
		t.Execute(w, p)
	}
}

func getSQLData() PatientInfoResponse {
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	if err != nil {
		panic(err)
	}
	sqlFile, err := os.ReadFile("./sql/GetValidatePatientInfo.sql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow(string(sqlFile), "Южно-Сахалинск", "Мира", 1)
	if err != nil {
		panic(err)
	}
	p := PatientInfoResponse{}
	err = row.Scan(&p.PatientID, &p.MOID, &p.MOOID, &p.MOName, &p.MOAddress, &p.MOPhone, &p.ResourceID, &p.ResourceName)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	return p
}
