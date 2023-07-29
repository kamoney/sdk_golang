package kamoney_sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/services/buy
func (s *sdk) ServicesBuy() (out kamoney_sdk_dtos.ServicesBuyRequestResponse, err error) {
	req, err := s.requestHandler.RequestHandler("GET", ENDPOINT_SERVICES_BUY, nil)
	if err != nil {
		log.Panicln("ServicesBuy 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("ServicesBuy 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("ServicesBuy 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("ServicesBuy 04: ", err.Error())
		return
	}

	return
}
