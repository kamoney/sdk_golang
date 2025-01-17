package public_kamoney

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"hash"
	"log"
	"net/http"
)

var (
	BASE_URL = "https://api2.kamoney.com.br/v2/"
)

type RequestHandler struct {
	PublicKey string
	SecretKey string
}

func (r *RequestHandler) SignRequest(req *http.Request) {
	var sig hash.Hash

	sig = hmac.New(sha256.New, []byte(r.SecretKey))
	sig.Write([]byte(req.URL.RawQuery))

	req.Header.Add("sign", hex.EncodeToString(sig.Sum(nil)))
}

func (r *RequestHandler) RequestHandler(method string, endpoint string, requestBody any) (*http.Request, error) {
	var req *http.Request
	var err error

	switch method {
	case "POST":
		body, err := json.Marshal(requestBody)
		if err != nil {
			log.Panicln("rH 00: ", err)
			return req, err
		}

		req, err = http.NewRequest(http.MethodPost, BASE_URL+endpoint, bytes.NewBuffer(body))
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}
	case "GET":
		req, err = http.NewRequest(http.MethodGet, BASE_URL+endpoint, nil)
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}

	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("public", r.PublicKey)

	return req, err
}
