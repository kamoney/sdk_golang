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

func (s *privateRequests) GetAPISecret(in kamoney_sdk_dtos.GetAPISecretRequestParams) (out kamoney_sdk_dtos.GetAPISecretRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("POST", ENDPOINT_SECURITY_API_SECRET(in.ID), in)
	if err != nil {
		log.Panicln("GAS 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce":    in.Nonce,
		"password": in.Password,
		"tfa":      fmt.Sprint(in.Tfa),
		"id":       fmt.Sprint(in.ID),
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("GAS 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("GAS 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("GAS 04: ", err.Error())
		return
	}

	return
}
