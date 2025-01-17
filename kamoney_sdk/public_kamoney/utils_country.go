package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/country
func (s *publicRequests) UtilsCountry() (out kamoney_sdk_dtos.UtilsCountryRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_UTILS_COUNTRY, nil)
	if err != nil {
		log.Panicln("UtilsCountry 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsCountry 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsCountry 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsCountry 04: ", err.Error())
		return
	}

	return
}
