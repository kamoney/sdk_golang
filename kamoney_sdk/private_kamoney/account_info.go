package private_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func (s *privateRequests) AccountInfo(in kamoney_sdk_dtos.AccountInfoRequestParams) (out kamoney_sdk_dtos.AccountInfoRequestResponse, err error) {
	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_INFO, in)
	if err != nil {
		log.Panicln("Active 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"name":          in.Name,
		"personal_id":   in.PersonalID,
		"date_of_birth": in.DateOfBirth,
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("Active 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Active 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Active 04: ", err.Error())
		return
	}

	return
}
