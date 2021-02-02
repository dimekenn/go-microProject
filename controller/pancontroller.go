package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/unistack-org/micro/model"
	"io/ioutil"
	"net/http"
)

var keys = []string{
	"5166546A576E5A7234753778214125442A472D4A614E645267556B5870327335",
	"3777217A24432646294A404E635266556A586E3272357538782F413F442A472D",
	"5367566B5970337336763979244226452948404D6251655468576D5A71347437",
	"645266556A586E3272357538782F413F4428472B4B6250655368566B59703373",
}

var PanToCustomer = func(res http.ResponseWriter, req *http.Request) {
	var (
		requestModel    model.Request
		wayFourResponse model.UFXmsgResponse
	)

	//Visa Request...
	err := json.NewDecoder(req.Body).Decode(&requestModel)
	if err != nil {
		panic(err)
		return
	}
	ibans, err := EncodePan2(requestModel)
	if err != nil {
		panic(err)
		return
	}
	err2 := Response(res, ibans)
	if err2 != nil {
		panic(err2)
	}

	//xml...
	msg := model.UFXmsg{
		Direction: "Rq",
		MsgType:   "Information",
		Version:   "2.3.80",
		Xmlns:     "http://www.w3.org/2001/XMLSchema-instance",
		MsgId:     1,
	}
	msgData := model.MsgData{}
	msg.Source = model.Source{App: "W4P"}
	msg.MsgData = msgData
	msg.MsgData.RefContractNumber = ibans

	//Way4 request
	url := "https://httpbin.org/post"
	xmlValue, err := xml.Marshal(msg)
	if err != nil {
		fmt.Println("error with xml marshal")
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlValue))
	if err != nil {
		fmt.Println("bad request")
	}
	request.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error with client do: s%", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	err = xml.Unmarshal(data, &wayFourResponse)
	if err != nil {
		fmt.Printf("error to get parse xml: %s", err)
	}

}

func Response(res http.ResponseWriter, response interface{}) error {
	res.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(res).Encode(response)
	return nil
}

func EncodePan2(request model.Request) ([]string, error) {
	var ibans []string
	for i := range request.Pan {
		iban, err := ExampleNewCBCDecrypter(keys[request.KeyId], request.Pan[i])
		if err != nil {
			return ibans, err
		}
		ibans = append(ibans, iban)
	}
	return ibans, nil
}

func ExampleNewCBCDecrypter(key string, s string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.

	decodeKey, _ := hex.DecodeString(key)
	ciphertext, _ := hex.DecodeString(s)
	block, err := aes.NewCipher(decodeKey)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		fmt.Println("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	substring := ciphertext[12:]
	fmt.Printf("%s\n", substring)

	return string(ciphertext), nil
}
