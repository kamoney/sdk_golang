package private_kamoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) DeleteAPI(in kamoney_sdk_dtos.DeleteAPIRequestParams) (out kamoney_sdk_dtos.DeleteAPIRequestResponse, err error) {
	req, err := s.r.RequestHandler("DELETE", ENDPOINT_SECURITY_API_DELETE(in.ID), in)
	if err != nil {
		log.Panicln("DA 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("DA 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("DA 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("DA 04: ", err.Error())
		return
	}

	return
}
