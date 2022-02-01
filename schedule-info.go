package main

import "net/http"

func GetHouseCallScheduleInfo(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "GetHouseCallScheduleInfo" {
		sendData("GetHouseCallScheduleInfo: " + action, w)
	}
}