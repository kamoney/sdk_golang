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

func (s *privateRequests) GetWithdrawInfo(in kamoney_sdk_dtos.GetWithdrawInfoRequestParams) (out kamoney_sdk_dtos.GetWithdrawInfoRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("GET", ENDPOINT_WITHDRAW_INFO(in.ID), in)
	if err != nil {
		log.Panicln("GWI 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce": in.Nonce,
		"id":    in.ID,
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("GWI 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("GWI 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GWI 04: ", err.Error())
		return
	}

	return
}
