package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/contact
func (s *sdk) UtilsContact(in kamoney_sdk_dtos.UtilsContactRequestParams) (out kamoney_sdk_dtos.UtilsContactRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("POST", ENDPOINT_UTILS_CONTACT, in)
	if err != nil {
		log.Panicln("UtilsContact 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsContact 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsContact 03: ", err.Error())
		return
	}
	// fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsContact 04: ", err.Error())
		return
	}

	return
}
