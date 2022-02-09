package main

import (
	"fmt"
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
		scheduleResponse := ScheduleInfoResponse {
				SessionID: "123123123123123123",
				Schedule: Schedule{
					Slots: []Slot{
						{SlotID:"id01", VisitTime:"time01", Duration:120},
						{SlotID:"id02", VisitTime:"time02", Duration:120},
						{SlotID:"id03", VisitTime:"time03", Duration:120},
						{SlotID:"id04", VisitTime:"time04", Duration:120},
				},
			},
		}
		fmt.Println(scheduleResponse)
		data, err := os.ReadFile(templateFile)
		if err != nil {
			panic(err)
		}
		t := template.New("tmp")
		t.Parse(string(data))
		w.Header().Add("Content-Type", "application/xml")
		t.Execute(w, scheduleResponse)
	}

}
