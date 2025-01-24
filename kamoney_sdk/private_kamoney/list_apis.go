package private_kamoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) ListAPIs(in kamoney_sdk_dtos.ListAPIsRequestParams) (out kamoney_sdk_dtos.ListAPIsRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_SECURITY_API, in)
	if err != nil {
		log.Panicln("CA 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("CA 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("CA 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("CA 04: ", err.Error())
		return
	}

	return
}
