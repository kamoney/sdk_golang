package private_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) AccountContact(in kamoney_sdk_dtos.AccountContactRequestParams) (out kamoney_sdk_dtos.AccountContactRequestResponse, err error) {
	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_CONTACT, in)
	if err != nil {
		log.Panicln("AC 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("AC 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AC 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AC 04: ", err.Error())
		return
	}

	return
}
