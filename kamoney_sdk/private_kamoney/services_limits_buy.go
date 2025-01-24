package private_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) GetServicesLimitsBuy(in kamoney_sdk_dtos.GetServicesLimitsBuyRequestParams) (out kamoney_sdk_dtos.GetServicesLimitsBuyRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_SERVICES_LIMITS_BUY, in)
	if err != nil {
		log.Panicln("GSLB 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("GSLB 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("GSLB 03: ", err.Error())
		return
	}
	// fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GSLB 04: ", err.Error())
		return
	}

	return
}
