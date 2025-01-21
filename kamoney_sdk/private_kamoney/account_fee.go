package private_kamoney

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func (s *privateRequests) GetFee(in kamoney_sdk_dtos.GetFeeRequestParams) (out kamoney_sdk_dtos.GetFeeRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("GET", ENDPOINT_SERVICES_FEE, in)
	if err != nil {
		log.Panicln("GSLB 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce": in.Nonce,
	})

	req.URL.RawQuery = queryStr
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
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GSLB 04: ", err.Error())
		return
	}

	return
}
