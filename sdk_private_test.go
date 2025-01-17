package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func TestAccountInfo(t *testing.T) {
	response, err := private.AccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams{
		Name:        "Igor Araújo da Silva (Dev Kamoney)",
		PersonalID:  "088.301.216-22",
		DateOfBirth: "05-08-1995",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountLocality(t *testing.T) {
	response, err := private.AccountLocality(kamoney_sdk_dtos.AccountLocalityRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
