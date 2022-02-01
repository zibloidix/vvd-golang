package main

import "net/http"

func CreateHouseCall(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "CreateHouseCall" {
		sendData("CreateHouseCall: "+action, w)
	}
}
