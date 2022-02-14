package main

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"text/template"
)

type CreateHouseCallResponse struct {
	RmisID string
	Comment string
	SlotID string
	VisitTime string
	Duration int
}

func CreateHouseCall(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "CreateHouseCall" {
		templateFile := respBase + action + "/" + respOk
		envelope, err := ParseSoapEnvelope(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + respErr01
		}
		createHouseCall, err := getSQLDataCreateHouseCall(envelope)
		if err != nil {
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
		t.Execute(w, createHouseCall)
	}

}

func getSQLDataCreateHouseCall(e Envelope) (CreateHouseCallResponse, error) {
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	c := CreateHouseCallResponse{}
	if err != nil {
		return c, err
	}
	defer db.Close()
	sqlFile, err := os.ReadFile("./sql/CreateHouseCall.sql")
	if err != nil {
		return c, err
	}
	slotId := getParamsFromCreateHouseCallRequest(e)
	row := db.QueryRow(string(sqlFile), slotId)
	err = row.Scan(&c.RmisID, &c.Comment, &c.SlotID, &c.VisitTime, &c.Duration)
	if err != nil {
		return c, err
	}
	return c, nil
}

func getParamsFromCreateHouseCallRequest(e Envelope) string {
	return e.Body.CreateHouseCallRequest.Slot_Id
}

