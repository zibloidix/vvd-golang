package main

import (
	"database/sql"
	"fmt"
	"github.com/aymerick/raymond"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"strconv"
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
		ctx := getCtx(p)
		fmt.Println(ctx)
		render, _ := raymond.Render(string(data), ctx)
		sendData(render, w)
	}
}

func getSQLData() PatientInfoResponse {
	db, err := sql.Open("sqlite3", "./db/vvd.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow("select 1001 PatientID, h.id   MOID, h.oid  MOOID, h.name  MOName, h.address MOAddress, h.phone  MOPhone, d.snils || '.' || d.doctor_code  ResourceID, d.name  || ' (' || s.name  || ')'  ResourceName from hospitals h join doctors d  on d.hospital_id = h.id join specs s  on d.spec_id = s.id join districts ds  on ds.hospital_id = h.id where 1=1 and ds.city = 'Южно-Сахалинск' and ds.street like '%Мира%' and ds.house = 1;")
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

func getCtx(p PatientInfoResponse) map[string]string {
	ctx := map[string]string{
		"PatientID":    strconv.Itoa(p.PatientID),
		"MOID":         strconv.Itoa(p.MOID),
		"MOOID":        p.MOOID,
		"MOName":       p.MOName,
		"MOAddress":    p.MOAddress,
		"MOPhone":      p.MOPhone,
		"ResourceID":   p.ResourceID,
		"ResourceName": p.ResourceName,
	}
	return ctx
}
