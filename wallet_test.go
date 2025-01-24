package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func TestListWallet(t *testing.T) {
	response, err := private.ListWallet(kamoney_sdk_dtos.ListWalletRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Filters is not working
func TestWalletExtract(t *testing.T) {
	response, err := private.WalletExtract(kamoney_sdk_dtos.WalletExtractRequestParams{
		Search: "R$",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestCreateWithdraw(t *testing.T) {
	response, err := private.CreateWithdraw(kamoney_sdk_dtos.CreateWithdrawRequestParams{
		Asset:  "R$",
		Type:   "CPF",
		Key:    "08830121622",
		Amount: 10,
		Tfa:    "123456",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestListWithdraw(t *testing.T) {
	response, err := private.ListWithdraw(kamoney_sdk_dtos.ListWithdrawRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetWithdrawInfo(t *testing.T) {
	response, err := private.GetWithdrawInfo(kamoney_sdk_dtos.GetWithdrawInfoRequestParams{
		ID: "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetWithdrawReceipt(t *testing.T) {
	response, err := private.GetWithdrawReceipt(kamoney_sdk_dtos.GetWithdrawReceiptRequestParams{
		ID: "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// With no data return array, and if theres data ? will return []struct or a struct ? how to handle ?
func TestGetWithdrawReceiptDownload(t *testing.T) {
	response, err := private.GetWithdrawReceiptDownload(kamoney_sdk_dtos.GetWithdrawReceiptDownloadRequestParams{
		ID:       "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H",
		FileName: "",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
