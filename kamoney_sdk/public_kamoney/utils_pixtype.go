package public_kamoney

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

// public/pixtype
func (s *publicRequests) UtilsPixType() (out kamoney_sdk_dtos.UtilsPixTypeRequestResponse, err error) {
	req, err := s.r.RequestHandler("GET", ENDPOINT_UTILS_PIX_TYPE, nil)
	if err != nil {
		log.Panicln("UtilsPixType 01: ", err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("UtilsPixType 02: ", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("UtilsPixType 03: ", err.Error())
		return
	}
	// fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("UtilsPixType 04: ", err.Error())
		return
	}

	return
}
