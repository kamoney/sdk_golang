package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/register
func (s *sdk) AccountRegister(in kamoney_sdk_dtos.AccountRegisterRequestParams) (out kamoney_sdk_dtos.AccountRegisterRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("POST", ENDPOINT_ACCOUNT_REGISTER, in)
	if err != nil {
		log.Panicln("Register 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("Register 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Register 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Register 04: ", err.Error())
		return
	}

	return
}
