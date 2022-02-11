package main

import (
	"fmt"
	"strings"
	"database/sql"
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
		envelope, err := ParseSoapEnvelope(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + respErr02
		}
		scheduleResponse, err := getSQLDataSchedule(envelope)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr01
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			templateFile = respBase + action + "/" + respErr02
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

func getSQLDataSchedule(e Envelope) (ScheduleInfoResponse, error) {
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	s := ScheduleInfoResponse{}
	if err != nil {
		return s, err
	}
	defer db.Close()
	sqlFile, err := os.ReadFile("./sql/GetHouseCallScheduleInfo.sql")
	if err != nil {
		return s, err
	}
	resId, startDate, endDate := getParamsFromScheduleRequest(e)
	snils, code := getSnilsAndCode(resId)
	fmt.Println(resId, startDate, endDate, snils, code)
	rows, err := db.Query(string(sqlFile), snils, code, startDate, endDate)
	if err != nil {
		return s, err
	}
	for rows.Next() {
		slot := new(Slot)
		err = rows.Scan(&slot.SlotID, &slot.VisitTime, &slot.Duration)
		
		fmt.Println(slot, err)
		if err != nil {
			return s, err
		}
		s.Schedule.Slots = append(s.Schedule.Slots, *slot)
	}
	fmt.Println(s)
	return s, nil
}

func getParamsFromScheduleRequest(e Envelope) (string, string, string) {
	resId := e.Body.GetHouseCallScheduleInfoRequest.Resource_Id
	startDate := e.Body.GetHouseCallScheduleInfoRequest.StartDateRange
	endDate := e.Body.GetHouseCallScheduleInfoRequest.EndDateRange
	return resId, startDate, endDate
}
func getSnilsAndCode(resId string) (string, string) {
	snils := ""
	code := ""
	n := strings.Index(resId, ".")
	snils = resId[:n]
	code = resId[n+1:]
	return snils, code
}