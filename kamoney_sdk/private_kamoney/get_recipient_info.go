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

func (s *privateRequests) GetRecipientInfo(in kamoney_sdk_dtos.GetRecipientInfoRequestParams) (out kamoney_sdk_dtos.GetRecipientInfoRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_RECIPIENTS_ID(in.ID), in)
	if err != nil {
		log.Panicln("CR 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce":           in.Nonce,
		"account_type":    fmt.Sprint(in.AccountType),
		"bank":            fmt.Sprint(in.Bank),
		"agency":          in.Agency,
		"account_number":  in.AccountNumber,
		"owner":           in.Owner,
		"personalid":      in.PersonalID,
		"description":     in.Description,
		"main":            fmt.Sprint(in.Main),
		"ddd":             in.DDD,
		"mobile_operator": in.MobileOperator,
		"phone_number":    in.PhoneNumber,
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("CR 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("CR 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("CR 04: ", err.Error())
		return
	}

	return
}
