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

func (s *privateRequests) ChangeRecipients(in kamoney_sdk_dtos.ChangeRecipientsRequestParams, id string) (out kamoney_sdk_dtos.ChangeRecipientsRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())
	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_RECIPIENTS_ID(id), in)
	if err != nil {
		log.Panicln("CR 01: ", err.Error())
		return
	}

	client := &http.Client{}

	q := s.mapToURLValues(s.gerQueryString(in))
	req.URL.RawQuery = q.Encode()
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
