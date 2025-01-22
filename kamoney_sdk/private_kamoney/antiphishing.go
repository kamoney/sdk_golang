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

func (s *privateRequests) AntiPhishing(in kamoney_sdk_dtos.AntiPhishingRequestParams) (out kamoney_sdk_dtos.AntiPhishingRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("POST", ENDPOINT_SECURITY_ANTIPHISHING, in)
	if err != nil {
		log.Panicln("AP 01: ", err.Error())
		return
	}

	client := &http.Client{}

	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"nonce":    in.Nonce,
		"tfa":      fmt.Sprint(in.Tfa),
		"password": in.Password,
		"phrase":   in.Phrase,
	})

	req.URL.RawQuery = queryStr
	s.r.signRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("AP 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AP 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AP 04: ", err.Error())
		return
	}

	return
}
