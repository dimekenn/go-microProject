package model

import "encoding/xml"

type UFXmsg struct {
	XMLName   xml.Name `xml:"UFXmsg"`
	Direction string   `xml:"direction,attr"`
	MsgType   string   `xml:"msg_type,attr"`
	Version   string   `xml:"version,attr"`
	Xsi       string   `xml:"xsi NoNameSpaceSchemaLocation,attr"`
	Xmlns     string   `xml:"xmlns xsi,attr"`
	MsgId     string   `xml:"MsgId"`
	Source    Source   `xml:"Source"`
	MsgData   MsgData  `xml:"MsgData"`
}

type Source struct {
	App string `xml:"app,attr"`
}

type MsgData struct {
	Text        string      `xml:",chardata"`
	Information Information `xml:"Information"`
}

type Information struct {
	Text        string    `xml:",chardata"`
	Institution string    `xml:"Institution"`
	ObjectType  string    `xml:"ObjectType"`
	ActionType  string    `xml:"ActionType"`
	ObjectFor   ObjectFor `xml:"ObjectFor"`
}

type ObjectFor struct {
	Text      string    `xml:",chardata"`
	ClientIDT ClientIDT `xml:"ClientIdt"`
}

type ClientIDT struct {
	Text              string `xml:",chardata"`
	RefContractNumber string `xml:"RefContractNumber"`
}
