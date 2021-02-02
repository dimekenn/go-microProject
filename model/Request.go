package model

type Request struct {
	KeyId int      `json:"keyId"`
	Pan   []string `json:"pan"`
}
