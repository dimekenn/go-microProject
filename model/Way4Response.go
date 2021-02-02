package model

import (
	"encoding/xml"
)

type UFXmsgResponse struct {
	XMLName   xml.Name `xml:"UFXmsg"`
	Direction string   `xml:"direction,attr"`
	MsgType   string   `xml:"msg_type,attr"`
	Version   string   `xml:"version,attr"`
	Xsi       string   `xml:"xsi noNameSpaceSchemaLocation,attr"`
	RespClass string   `xml:"resp_class,attr"`
	RespCode  string   `xml:"resp_code,code"`
	RespText  string   `xml:"resp_text,attr"`
	Xmlns     string   `xml:"xmlns xsi,attr"`
	MsgId     string   `xml:"MsgId"`
	Source    Source2  `xml:"Source"`
	MsgData   MsgData  `xml:"MsgData"`
}

type Source2 struct {
	App string `xml:"app,attr"`
}

type MsgData2 struct {
	Text        string      `xml:",chardata"`
	Information Information `xml:"Information"`
}

type Information2 struct {
	Text        string     `xml:",chardata"`
	Institution string     `xml:"Institution"`
	ObjectType  string     `xml:"ObjectType"`
	ActionType  string     `xml:"ActionType"`
	ObjectFor   ObjectFor2 `xml:"ObjectFor"`
	DataRs      DataRs     `xml:"DataRs"`
}

type ObjectFor2 struct {
	Text      string     `xml:",chardata"`
	ClientIDT ClientIDT2 `xml:"ClientIdt"`
}

type ClientIDT2 struct {
	Text              string `xml:",chardata"`
	RefContractNumber string `xml:"RefContractNumber"`
}

type DataRs struct {
	ClientRs ClientRs `xml:"ClientRs"`
}

type ClientRs struct {
	Client Client `xml:"Client"`
}

type Client struct {
	OrderDrt       string      `xml:"OrderDrt"`
	ClientType     string      `xml:"ClientType"`
	ClientStrategy string      `xml:"ClientStrategy"`
	ClientInfo     ClientInfo  `xml:"ClientInfo"`
	PlasticInfo    PlasticInfo `xml:"PlasticInfo"`
	DateOpen       string      `xml:"DateOpen"`
	BaseAddress    BaseAddress `xml:"BaseAddress"`
	AddInfo        AddInfo     `xml:"AddInfo"`
}

type ClientInfo struct {
	ClientNumber  string `xml:"ClientNumber"`
	RegNumberType string `xml:"RegNumberType"`
	ShortNumber   string `xml:"ShortNumber"`
	LastName      string `xml:"LastName"`
	Country       string `xml:"Country"`
	Language      string `xml:"Language"`
	CompanyName   string `xml:"CompanyName"`
}

type PlasticInfo struct {
	FirstName string `xml:"FirstName"`
	LastName  string `xml:"LastName"`
}

type BaseAddress struct {
	State       string   `xml:"State"`
	City        string   `xml:"City"`
	PostalCode  string   `xml:"PostalCode"`
	AddressLine []string `xml:"AddressLine"`
}

type AddInfo struct {
	AddInfo []string `xml:"AddInfo"`
	ExtraRs string   `xml:"ExtraRs"`
}
