package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/recovery/confirm
func (s *publicRequests) AccountRecoveryConfirm(in kamoney_sdk_dtos.AccountRecoveryConfirmRequestParams) (out kamoney_sdk_dtos.AccountRecoveryConfirmRequestResponse, err error) {
	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_RECOVERY_CONFIRM, in)
	if err != nil {
		log.Panicln("Recovery 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("Recovery 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Recovery 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Recovery 04: ", err.Error())
		return
	}

	return
}
