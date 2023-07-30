package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/active
func (s *sdk) AccountActive(in kamoney_sdk_dtos.AccountActiveRequestParams) (out kamoney_sdk_dtos.AccountActiveRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("POST", ENDPOINT_ACCOUNT_ACTIVE, in)
	if err != nil {
		log.Panicln("Active 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("Active 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Active 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Active 04: ", err.Error())
		return
	}

	return
}
