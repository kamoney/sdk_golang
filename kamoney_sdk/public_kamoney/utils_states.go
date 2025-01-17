package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/country/{country_id}/state
func (s *publicRequests) UtilsState(country_id int64) (out kamoney_sdk_dtos.UtilsStateRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_UTILS_COUNTRY_STATE(country_id), nil)
	if err != nil {
		log.Panicln("UtilsState 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsState 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsState 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsState 04: ", err.Error())
		return
	}

	return
}
