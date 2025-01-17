package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// Unecessary endpoint.
func (s *publicRequests) Auth(in kamoney_sdk_dtos.AuthRequestParams) (out kamoney_sdk_dtos.AuthRequestResponse, err error) {
	req, err := s.r.RequestHandler("POST", ENDPOINT_AUTH, in)
	if err != nil {
		log.Panicln("A 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("A 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("A 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("A 04: ", err.Error())
		return
	}

	return out, err
}
