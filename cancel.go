package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"text/template"
)

type CancelHouseCallResponse struct {
	RmisID string
}

func CancelHouseCall(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "CancelHouseCall" {

		templateFile := respBase + action + "/" + respOk
		envelope, err := ParseSoapEnvelope(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr02
		}
		cancelHouse, err := getSQLDataCancelHouseCall(envelope)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr02
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr01
		}
		data, err := os.ReadFile(templateFile)
		if err != nil {
			panic(err)
		}
		t := template.New("tmp")
		t.Parse(string(data))
		w.Header().Add("Content-Type", "application/xml")
		t.Execute(w, cancelHouse)
	}
}

func getSQLDataCancelHouseCall(e Envelope) (CancelHouseCallResponse, error) {
	c := CancelHouseCallResponse{}
	c.RmisID = e.Body.CancelHouseCallRequest.HC_Id_Rmis
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	if err != nil {
		return c, err
	}
	defer db.Close()
	sqlFile, err := os.ReadFile("./sql/CancelHouseCall.sql")
	if err != nil {
		return c, err
	}
	row := db.QueryRow(string(sqlFile), c.RmisID)
	err = row.Scan(&c.RmisID)
	if err != nil {
		return c, err
	}
	return c, nil
}
