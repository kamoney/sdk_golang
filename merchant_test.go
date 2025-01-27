package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestCreateMerchant(t *testing.T) {
	response, err := private.CreateMerchant(kamoney_sdk_dtos.CreateMerchantRequestParams{
		Nonce:   fmt.Sprint(utility.GenNonce()),
		Asset:   "LTC",
		Network: "LTC",
		Amount:  100,
		Email:   "igorasft@gmail.com",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListMerchant(t *testing.T) {
	response, err := private.ListMerchant(kamoney_sdk_dtos.ListMerchantRequestParams{
		Begin:  "",
		End:    "",
		Status: "PENDING",
		Search: "",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetMerchantInfo(t *testing.T) {
	response, err := private.GetMerchantInfo(kamoney_sdk_dtos.GetMerchantInfoRequestParams{}, "MKM84797676")
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
