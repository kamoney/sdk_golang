package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/services/order
func (s *publicRequests) ServicesOrder() (out kamoney_sdk_dtos.ServicesOrderRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_SERVICES_ORDER, nil)
	if err != nil {
		log.Panicln("ServicesOrder 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("ServicesOrder 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("ServicesOrder 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("ServicesOrder 04: ", err.Error())
		return
	}

	return
}
