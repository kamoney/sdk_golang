package kamoney_sdk_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func printStruct(s interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)
	fmt.Println(s)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := fieldType.Name

		// fmt.Println(fieldType.Name, field.Kind())
		if field.Kind() == reflect.Struct {
			// Chamada recursiva para structs aninhadas
			nestedFields := printStruct(field.Interface())
			for _, v := range nestedFields {
				// Adiciona as chaves da struct aninhada no map atual
				result[strings.ToLower(fieldName)] = v
			}
		} else if field.Kind() == reflect.Slice {
			// Se o campo for um slice, iteramos sobre os elementos
			for j := 0; j < field.Len(); j++ {
				elem := field.Index(j).Interface()
				nestedFields := printStruct(elem)
				for k, v := range nestedFields {
					result[fmt.Sprintf("%s%s", strings.ToLower(fieldName), k)] = v
				}
			}
		} else {
			// Para campos simples (não structs nem slices)
			result[strings.ToLower(fieldName)] = fmt.Sprintf("%v", field.Interface())
		}
	}

	return result
}

type OrderRequest struct {
	Asset        string `json:"asset"`
	Network      string `json:"network"`
	PaymentSlips []struct {
		Barcode     string  `json:"barcode"`
		Institution string  `json:"institution"`
		Amount      float64 `json:"amount"`
		DueDate     string  `json:"due_date"`
	} `json:"payment_slips"`
}

var obj OrderRequest

func TestGetString(t *testing.T) {
	obj = OrderRequest{
		Asset:   "LTC",
		Network: "LTC",
		PaymentSlips: []struct {
			Barcode     string  `json:"barcode"`
			Institution string  `json:"institution"`
			Amount      float64 `json:"amount"`
			DueDate     string  `json:"due_date"`
		}{
			{
				Barcode:     "23793380296101222882356006333308191510000002000",
				Institution: "Banco NuBank",
				Amount:      100.50,
				DueDate:     "2023-12-31",
			},
		},
	}

	result := printStruct(obj)
	fmt.Println(result)
}
