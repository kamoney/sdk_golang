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
		Status: "CANCELED",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetBuyInfo(t *testing.T) {
	response, err := private.GetBuyInfo(kamoney_sdk_dtos.GetBuyInfoRequestParams{}, "BKM81604797")

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetBuyNewQRCode(t *testing.T) {
	response, err := private.GetBuyNewQRCode(kamoney_sdk_dtos.GetBuyNewQRCodeRequestParams{}, "BKM81604797")

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetBuyPrivateKey(t *testing.T) {
	response, err := private.GetBuyPrivateKey(kamoney_sdk_dtos.GetBuyPrivateKeyRequestParams{
		Password: "",
		TFA:      "",
	}, "BKM81604797")

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestCreateBuy(t *testing.T) {
	response, err := private.CreateBuy(kamoney_sdk_dtos.CreateBuyRequestParams{
		Nonce:         fmt.Sprint(utility.GenNonce()),
		Asset:         "BTC",
		Network:       "BTC",
		Amount:        500,
		WalletType:    1,
		PaymentMethod: "pix",
		Address:       "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
		BaseCoin:      "",
		Quantity:      0,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
