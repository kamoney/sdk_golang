package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/country/{country_id}/state/{state_id}/city
func (s *sdk) UtilsCity(country_id, state_id int64) (out kamoney_sdk_dtos.UtilsCityRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("GET", ENDPOINT_UTILS_COUNTRY_CITY(country_id, state_id), nil)
	if err != nil {
		log.Panicln("UtilsCity 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsCity 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsCity 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsCity 04: ", err.Error())
		return
	}

	return
}
