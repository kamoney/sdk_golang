package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/currency
func (s *publicRequests) UtilsCurrency() (out kamoney_sdk_dtos.UtilsCurrencyRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_UTILS_CURRENCY, nil)
	if err != nil {
		log.Panicln("Currency 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("Currency 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Currency 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Currency 04: ", err.Error())
		return
	}

	return
}
