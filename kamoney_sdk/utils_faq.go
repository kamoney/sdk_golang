package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/faq
func (s *sdk) UtilsFaq() (out kamoney_sdk_dtos.UtilsFaqRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("GET", ENDPOINT_UTILS_BANK, nil)
	if err != nil {
		log.Panicln("UtilsFaq 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsFaq 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsFaq 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsFaq 04: ", err.Error())
		return
	}

	return
}
