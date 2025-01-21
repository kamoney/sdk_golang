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

// account/locality
func (s *privateRequests) GetAccountLocality(in kamoney_sdk_dtos.GetAccountLocalityRequestParams) (out kamoney_sdk_dtos.GetAccountLocalityRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_LOCALITY, in)
	if err != nil {
		log.Panicln("Active 01: ", err.Error())
		return
	}

	client := &http.Client{}

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
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Active 04: ", err.Error())
		return
	}

	return
}

func (s *privateRequests) AccountLocality(in kamoney_sdk_dtos.AccountLocalityRequestParams) (out kamoney_sdk_dtos.AccountLocalityRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("POST", ENDPOINT_ACCOUNT_LOCALITY, in)
	if err != nil {
		log.Panicln("Active 01: ", err.Error())
		return
	}

	client := &http.Client{}
	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
		"zipcode":      in.Zipcode,
		"street":       in.Street,
		"number":       fmt.Sprint(in.Number),
		"complement":   in.Complement,
		"neighborhood": in.Neighborhood,
		"city":         fmt.Sprint(in.City),
		"state":        fmt.Sprint(in.State),
		"nonce":        in.Nonce,
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
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("Active 04: ", err.Error())
		return
	}

	return
}
