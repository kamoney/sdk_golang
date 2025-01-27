package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestListBuy(t *testing.T) {
	response, err := private.ListBuy(kamoney_sdk_dtos.ListBuyRequestParams{
		Begin:  "",
		End:    "",
		Search: "",
		Status: "PENDING",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetBuyInfo(t *testing.T) {
	response, err := private.GetBuyInfo(kamoney_sdk_dtos.GetBuyInfoRequestParams{}, "BKM18650087")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

// func TestGetBuyNewQRCode(t *testing.T) {
// 	response, err := private.GetBuyNewQRCode(kamoney_sdk_dtos.GetBuyNewQRCodeRequestParams{}, "BKM18650087")

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Log(response)
// }

// func TestGetBuyPrivateKey(t *testing.T) {
// 	response, err := private.GetBuyPrivateKey(kamoney_sdk_dtos.GetBuyPrivateKeyRequestParams{
// 		Password: "",
// 		TFA:      "",
// 	}, "BKM18650087")

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Log(response)
// }

func TestCreateBuy(t *testing.T) {
	response, err := private.CreateBuy(kamoney_sdk_dtos.CreateBuyRequestParams{
		Nonce:         fmt.Sprint(utility.GenNonce()),
		Asset:         "BTC",
		Network:       "BTC",
		Amount:        2500.25,
		WalletType:    1,
		PaymentMethod: "pix",
		Address:       "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
		BaseCoin:      "",
		Quantity:      0,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
