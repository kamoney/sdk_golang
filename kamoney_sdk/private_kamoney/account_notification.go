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

func (s *privateRequests) GetAccountNotification(in kamoney_sdk_dtos.GetAccountNotificationRequestParams) (out kamoney_sdk_dtos.GetAccountNotificationRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("GET", ENDPOINT_ACCOUNT_NOTIFICATION, in)
	if err != nil {
		log.Panicln("AN 01: ", err.Error())
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
		log.Panicln("AN 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("AN 03: ", err.Error())
		return
	}
	// fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("AN 04: ", err.Error())
		return
	}

	return
}

func (s *privateRequests) UpdateAccountNotificationReadAll(in kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestParams) (out kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("PUT", ENDPOINT_ACCOUNT_NOTIFICATION, in)
	if err != nil {
		log.Panicln("UANRA 01: ", err.Error())
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
		log.Panicln("UANRA 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UANRA 03: ", err.Error())
		return
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UANRA 04: ", err.Error())
		return
	}

	return
}

func (s *privateRequests) UpdateAccountNotificationReadId(in kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestParams) (out kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestResponse, err error) {
	in.Nonce = fmt.Sprint(utility.GenNonce())

	req, err := s.r.RequestHandler("PUT", ENDPOINT_ACCOUNT_NOTIFICATION+"/"+in.ID, nil)
	if err != nil {
		log.Panicln("UANRA 01: ", err.Error())
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
		log.Panicln("UANRA 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UANRA 03: ", err.Error())
		return
	}

	fmt.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UANRA 04: ", err.Error())
		return
	}

	return
}
