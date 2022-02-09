package main

import (
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"text/template"
)

type ScheduleInfoResponse struct {
	SessionID string
	Schedule Schedule 
}

type Schedule struct {
	Slots []Slot
}

type Slot struct {
	SlotID string
	VisitTime string
	Duration int
}

func GetHouseCallScheduleInfo(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "GetHouseCallScheduleInfo" {
		templateFile := respBase + action + "/" + respOk
		//envelope, err := ParseSoapEnvelope(r.Body)
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	templateFile = respBase + respErr02
		//}
		slots := []Slot{
			{"id01", "time01", 120},
			{"id02", "time02", 120},
			{"id03", "time03", 120},
			{"id04", "time04", 120},
		}
		data, err := os.ReadFile(templateFile)
		if err != nil {
			panic(err)
		}
		t := template.New("tmp")
		t.Parse(string(data))
		w.Header().Add("Content-Type", "application/xml")
		t.Execute(w, slots)
	}

}
