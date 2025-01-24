package private_kamoney

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/kamoney/sdk_golang/utility"
)

var (
	BASE_URL = "https://api2.kamoney.com.br/v2/"
)

type RequestHandler struct {
	PublicKey string
	SecretKey string
}

func (s *privateRequests) mapToURLValues(m map[string]string) url.Values {
	values := url.Values{}
	for key, value := range m {
		values.Add(key, value) // Adiciona os pares chave-valor
	}
	return values
}

func (s *privateRequests) gerQueryString(obj interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	fmt.Println("Obj: ", obj)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		// fieldName := fieldType.Name

		// fmt.Println(fieldType.Name, field.Kind())
		if field.Kind() == reflect.Struct {
			// Chamada recursiva para structs aninhadas
			nestedFields := s.gerQueryString(field.Interface())
			for _, v := range nestedFields {
				// Adiciona as chaves da struct aninhada no map atual
				result[strings.ToLower(fieldType.Tag.Get("json"))] = v
			}
		} else if field.Kind() == reflect.Slice {
			// Se o campo for um slice, iteramos sobre os elementos
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j).Interface()
				nestedFields := s.gerQueryString(elem)
				for k, v := range nestedFields {
					result[fmt.Sprintf("%s%s", fieldType.Tag.Get("json"), k)] = v
				}
			}
		} else {
			// Para campos simples (não structs nem slices)
			result[fieldType.Tag.Get("json")] = fmt.Sprintf("%v", field.Interface())
		}
	}

	result["nonce"] = fmt.Sprint(utility.GenNonce())
	fmt.Println(result)

	return result
}

// func (s *privateRequests) gerQueryString(q url.Values, values map[string]string) (queryStr string) {
// 	for k, v := range values {
// 		q.Add(k, v)
// 	}

// 	queryStr = q.Encode()
// 	return
// }

func (r *RequestHandler) signRequest(req *http.Request) {
	var sig hash.Hash
	fmt.Println(req.URL.RawQuery)
	sig = hmac.New(sha512.New, []byte(r.SecretKey))
	sig.Write([]byte(req.URL.RawQuery))

	req.Header.Add("sign", hex.EncodeToString(sig.Sum(nil)))
}

func (r *RequestHandler) RequestHandler(method string, endpoint string, requestBody any) (*http.Request, error) {
	var req *http.Request
	var err error

	switch method {
	case "POST":
		body, err := json.Marshal(requestBody)
		if err != nil {
			log.Panicln("rH 00: ", err)
			return req, err
		}

		// It's necessary transform struct to map format
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		body, err = json.Marshal(result)
		if err != nil {
			log.Panicln("rH 00: ", err)
			return req, err
		}
		fmt.Println(string(body))
		req, err = http.NewRequest(http.MethodPost, BASE_URL+endpoint, bytes.NewBuffer(body))
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}
	case "DELETE":
		body, err := json.Marshal(requestBody)
		if err != nil {
			log.Panicln("rH 00: ", err)
			return req, err
		}

		// It's necessary transform struct to map format
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		body, err = json.Marshal(result)
		if err != nil {
			log.Panicln("rH 00: ", err)
			return req, err
		}
		fmt.Println(string(body))
		req, err = http.NewRequest(http.MethodDelete, BASE_URL+endpoint, bytes.NewBuffer(body))
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}
	case "GET":
		req, err = http.NewRequest(http.MethodGet, BASE_URL+endpoint, nil)
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}

	case "PUT":
		req, err = http.NewRequest(http.MethodPut, BASE_URL+endpoint, nil)
		if err != nil {
			log.Panicln("rH 01: ", err)
			return req, err
		}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("public", r.PublicKey)

	return req, err
}
