package private_kamoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) GetMerchantInfo(in kamoney_sdk_dtos.GetMerchantInfoRequestParams, id string) (out kamoney_sdk_dtos.GetMerchantInfoRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_MERCHANT_INFO(id), in)
	if err != nil {
		log.Panicln("LO 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("LO 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("LO 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("LO 04: ", err.Error())
		return
	}

	return
}
