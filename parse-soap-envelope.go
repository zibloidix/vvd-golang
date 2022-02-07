package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body
}

type Header struct {
	XMLName  xml.Name `xml:"Header"`
	Security Security
}

type Security struct {
	XMLName             xml.Name `xml:"Security"`
	BinarySecurityToken string   `xml:"BinarySecurityToken"`
	Signature           Signature
}

type Signature struct {
	XMLName        xml.Name `xml:"Signature"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        KeyInfo
}

type SignedInfo struct {
	XMLName                xml.Name `xml:"SignedInfo"`
	CanonicalizationMethod string   `xml:"CanonicalizationMethod"`
	SignatureMethod        string   `xml:"SignatureMethod"`
	Reference              Reference
}

type Reference struct {
	XMLName      xml.Name `xml:"Reference"`
	DigestMethod string   `xml:"DigestMethod"`
	DigestValue  string   `xml:"DigestValue"`
}

type Transforms struct {
	XMLName   xml.Name `xml:"Transforms"`
	Transform string   `xml:"Transform"`
}

type KeyInfo struct {
	XMLName                xml.Name `xml:"KeyInfo"`
	SecurityTokenReference SecurityTokenReference
}

type SecurityTokenReference struct {
	XMLName   xml.Name `xml:"SecurityTokenReference"`
	Reference string   `xml:"Reference"`
}

type Body struct {
	XMLName                         xml.Name `xml:"Body"`
	GetValidatePatientInfoRequest   GetValidatePatientInfoRequest
	GetHouseCallScheduleInfoRequest GetHouseCallScheduleInfoRequest
	CreateHouseCallRequest          CreateHouseCallRequest
	CancelHouseCallRequest          CancelHouseCallRequest
}

type PatientData struct {
	XMLName     xml.Name `xml:"Patient_Data"`
	OMS_Number  string   `xml:"OMS_Number"`
	OMS_Series  string   `xml:"OMS_Series"`
	SNILS       string   `xml:"SNILS"`
	First_Name  string   `xml:"First_Name"`
	Last_Name   string   `xml:"Last_Name"`
	Middle_Name string   `xml:"Middle_Name"`
	Birth_Date  string   `xml:"Birth_Date"`
	Sex         string   `xml:"Sex"`
}

type АpplicantData struct {
	XMLName      xml.Name `xml:"Applicant_Data"`
	First_Name   string   `xml:"First_Name"`
	Last_Name    string   `xml:"Last_Name"`
	Middle_Name  string   `xml:"Middle_Name"`
	SNILS        string   `xml:"SNILS"`
	Mobile_Phone string   `xml:"Mobile_Phone"`
	Email        string   `xml:"Email"`
}

type GetValidatePatientInfoRequest struct {
	XMLName               xml.Name `xml:"GetValidatePatientInfoRequest"`
	Session_ID            string   `xml:"Session_ID"`
	Patient_Data          PatientData
	ApplicantData         АpplicantData
	Cod_Kladr_Fias        string `xml:"Cod_Kladr_Fias"`
	Address_Str           string `xml:"Address_Str"`
	Adr_Region            string `xml:"Adr_Region"`
	Adr_Area              string `xml:"Adr_Area"`
	Adr_City              string `xml:"Adr_City"`
	Adr_City_Area         string `xml:"Adr_City_Area"`
	Adr_Place             string `xml:"Adr_Place"`
	Adr_Street            string `xml:"Adr_Street"`
	Adr_Additional_Area   string `xml:"Adr_Additional_Area"`
	Adr_Additional_Street string `xml:"Adr_Additional_Street"`
	Adr_House             string `xml:"Adr_House"`
	Adr_Housing           string `xml:"Adr_Housing"`
	Adr_Structure         string `xml:"Adr_Structure"`
	Adr_Apartment         string `xml:"Adr_Apartment"`
	Adr_Post_Index        string `xml:"Adr_Post_Index"`
	Reason_Task           string `xml:"Reason_Task"`
}

type GetHouseCallScheduleInfoRequest struct {
	XMLName        xml.Name `xml:"GetHouseCallScheduleInfoRequest"`
	Session_ID     string   `xml:"Session_ID"`
	Resource_Id    string   `xml:"Resource_Id"`
	StartDateRange string   `xml:"StartDateRange"`
	EndDateRange   string   `xml:"EndDateRange"`
}

type CreateHouseCallRequest struct {
	XMLName    xml.Name `xml:"CreateHouseCallRequest"`
	Session_ID string   `xml:"Session_ID"`
	Slot_Id    string   `xml:"Slot_Id"`
}

type CancelHouseCallRequest struct {
	XMLName    xml.Name `xml:"CancelHouseCallRequest"`
	HC_Id_Rmis string   `xml:"HC_Id_Rmis"`
}

func ParseSoapEnvelope(body io.ReadCloser) Envelope {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	var Envelope Envelope
	err = xml.Unmarshal(data, &Envelope)
	if err != nil {
		panic(err)
	}
	return Envelope
}
