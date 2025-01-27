package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestListWallet(t *testing.T) {
	response, err := private.ListWallet(kamoney_sdk_dtos.ListWalletRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

// Filters is not working
func TestWalletExtract(t *testing.T) {
	response, err := private.GetWalletExtract(kamoney_sdk_dtos.GetWalletExtractRequestParams{
		Search: "R$",
		Begin:  "",
		End:    "",
		Type:   "",
		Page:   0,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestCreateWithdraw(t *testing.T) {
	response, err := private.CreateWithdraw(kamoney_sdk_dtos.CreateWithdrawRequestParams{
		Asset:  "R$",
		Type:   "CPF",
		Key:    "08830121622",
		Amount: 10,
		Tfa:    "123456",
		Nonce:  fmt.Sprint(utility.GenNonce()),
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListWithdraw(t *testing.T) {
	response, err := private.ListWithdraw(kamoney_sdk_dtos.ListWithdrawRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetWithdrawInfo(t *testing.T) {
	response, err := private.GetWithdrawInfo(kamoney_sdk_dtos.GetWithdrawInfoRequestParams{}, "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetWithdrawReceipt(t *testing.T) {
	response, err := private.GetWithdrawReceipt(kamoney_sdk_dtos.GetWithdrawReceiptRequestParams{}, "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetWithdrawReceiptDownload(t *testing.T) {
	response, err := private.GetWithdrawReceiptDownload(kamoney_sdk_dtos.GetWithdrawReceiptDownloadRequestParams{
		FileName: "",
	}, "8WWrwErRURkCviDdO1E0CibEk7uC1gbpuLUkgT0Ga0H")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
