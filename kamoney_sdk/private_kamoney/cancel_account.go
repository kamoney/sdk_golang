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

func (s *privateRequests) CancelAccount(in kamoney_sdk_dtos.CancelAccountRequestParams) (out kamoney_sdk_dtos.CancelAccountRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("POST", ENDPOINT_SECURITY_ACTION, in)
	if err != nil {
		log.Panicln("CE 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce":    in.Nonce,
		"password": in.Password,
		"terms":    fmt.Sprint(in.Terms),
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("CE 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("CE 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("CE 04: ", err.Error())
		return
	}

	return
}
