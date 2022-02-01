package main

import "net/http"

func GetValidatePatientInfo(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("soapaction")
	if action == "GetValidatePatientInfo" {
		sendData("GetValidatePatientInfo: " + action, w)
	}
}