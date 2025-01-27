package kamoney_sdk_test

import (
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func TestCreatePaymentLink(t *testing.T) {
	response, err := private.CreatePaymentLink(kamoney_sdk_dtos.CreatePaymentLinkRequestParams{
		Label:  "Test",
		Amount: 100,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListPaymentLink(t *testing.T) {
	response, err := private.ListPaymentLink(kamoney_sdk_dtos.ListPaymentLinkRequestParams{
		Begin:  "",
		End:    "",
		Search: "",
		Status: "CANCELED",
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
	}, 843)

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
