package private_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) GetAccountInfo(in kamoney_sdk_dtos.ChangeAccountInfoRequestParams) (out kamoney_sdk_dtos.GetAccountInfoRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_INFO, nil)
	if err != nil {
		log.Panicln("AI 01: ", err.Error())
		return
	}

	client := &http.Client{}

	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("AI 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AI 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AI 04: ", err.Error())
		return
	}

	return
}
