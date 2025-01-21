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

func (s *privateRequests) AccountInfo(in kamoney_sdk_dtos.AccountInfoRequestParams) (out kamoney_sdk_dtos.AccountInfoRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_INFO, in)
	if err != nil {
		log.Panicln("AI 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"name":          in.Name,
		"personal_id":   in.PersonalID,
		"date_of_birth": in.DateOfBirth,
		"nonce":         in.Nonce,
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("AI 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AI 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AI 04: ", err.Error())
		return
	}

	return
}

func (s *privateRequests) GetAccountInfo(in kamoney_sdk_dtos.AccountInfoRequestParams) (out kamoney_sdk_dtos.GetAccountInfoRequestResponse, err error) {
	in.Nonce = fmt.Sprint(in.Nonce)

	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_INFO, nil)
	if err != nil {
		log.Panicln("AI 01: ", err.Error())
		return
	}

	client := &http.Client{}

	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("AI 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AI 03: ", err.Error())
		return
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AI 04: ", err.Error())
		return
	}

	return
}
