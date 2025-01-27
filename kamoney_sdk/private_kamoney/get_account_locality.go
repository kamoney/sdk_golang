package private_kamoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// account/locality
func (s *privateRequests) GetAccountLocality(in kamoney_sdk_dtos.GetAccountLocalityRequestParams) (out kamoney_sdk_dtos.GetAccountLocalityRequestResponse, err error) {

	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_LOCALITY, in)
	if err != nil {
		log.Panicln("GAL 01: ", err.Error())
		return
	}

	client := &http.Client{}

	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("GAL 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("GAL 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GAL 04: ", err.Error())
		return
	}

	return
}
