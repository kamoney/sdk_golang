package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestCreatePaymentLink(t *testing.T) {
	response, err := private.CreatePaymentLink(kamoney_sdk_dtos.CreatePaymentLinkRequestParams{
		Label:  "SDK Test",
		Amount: 100,
		Nonce:  fmt.Sprint(utility.GenNonce()),
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListPaymentLink(t *testing.T) {
	response, err := private.ListPaymentLink(kamoney_sdk_dtos.ListPaymentLinkRequestParams{
		Begin:  "2025-01-27",
		End:    "",
		Search: "",
		Status: "PENDING",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestDeletePaymentLink(t *testing.T) {
	response, err := private.DeletePaymentLink(kamoney_sdk_dtos.DeletePaymentLinkRequestParams{
		Begin:  "",
		End:    "",
		Search: "",
		Status: "",
	}, 845)

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
