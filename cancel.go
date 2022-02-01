package main

import "net/http"

func CancelHouseCall(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "CancelHouseCall" {
		sendData("CancelHouseCall: " + action, w)
	}
}