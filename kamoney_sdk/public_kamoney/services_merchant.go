package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/services/merchant
func (s *publicRequests) ServicesMerchant() (out kamoney_sdk_dtos.ServicesMerchantRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_SERVICES_MERCHANT, nil)
	if err != nil {
		log.Panicln("ServicesMerchant 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("ServicesMerchant 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("ServicesMerchant 03: ", err.Error())
		return
	}
	// fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("ServicesMerchant 04: ", err.Error())
		return
	}

	return
}
