package private_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) GetAccountKyc(in kamoney_sdk_dtos.GetAccountKycRequestParams) (out kamoney_sdk_dtos.GetAccountKycRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_KYC, in)
	if err != nil {
		log.Panicln("GAK 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("GAK 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("GAK 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GAK 04: ", err.Error())
		return
	}

	return
}
