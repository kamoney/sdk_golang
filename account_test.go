package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestGetAccountInfo(t *testing.T) {
	response, err := private.GetAccountInfo(kamoney_sdk_dtos.ChangeAccountInfoRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestChangeAccountInfo(t *testing.T) {
	response, err := private.ChangeAccountInfo(kamoney_sdk_dtos.ChangeAccountInfoRequestParams{
		Name:        "Igor Araújo da Silva",
		PersonalID:  "08830121622",
		DateOfBirth: "1995-08-05",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetAccountLocality(t *testing.T) {
	response, err := private.GetAccountLocality(kamoney_sdk_dtos.GetAccountLocalityRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestChangeAccountLocality(t *testing.T) {
	response, err := private.ChangeAccountLocality(kamoney_sdk_dtos.ChangeAccountLocalityRequestParams{
		Zipcode:      "35430232",
		Street:       "Av. Dom Bosco",
		Number:       490,
		Complement:   "201",
		Neighborhood: "Palmeiras",
		City:         3592,
		Nonce:        fmt.Sprint(utility.GenNonce()),
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

/*Needs Bearer Token*/
func TestChangeAccountContact(t *testing.T) {
	response, err := private.ChangeAccountContact(kamoney_sdk_dtos.ChangeAccountContactRequestParams{
		Whatsapp: "986621962",
		Telegram: "@igorasf",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetAccountHistory(t *testing.T) {
	response, err := private.GetAccountHistory(kamoney_sdk_dtos.GetAccountHistoryRequestParams{
		Page:   1,
		Search: "#OKM92006666",
		Date:   "",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
func TestGetAccountNotification(t *testing.T) {
	response, err := private.GetAccountNotification(kamoney_sdk_dtos.GetAccountNotificationRequestParams{
		Page: 1,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUpdateAccountNotificationReadAll(t *testing.T) {
	response, err := private.UpdateAccountNotificationReadAll(kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUpdateAccountNotificationReadId(t *testing.T) {
	response, err := private.UpdateAccountNotificationReadId(kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestParams{}, "2937")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetAccountKyc(t *testing.T) {
	response, err := private.GetAccountKyc(kamoney_sdk_dtos.GetAccountKycRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetServicesLimitsOrder(t *testing.T) {
	response, err := private.GetServicesLimitsOrder(kamoney_sdk_dtos.GetServicesLimitsOrderRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetServicesLimitsMerchant(t *testing.T) {
	response, err := private.GetServicesLimitsMerchant(kamoney_sdk_dtos.GetServicesLimitsMerchantRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetServicesLimitsBuy(t *testing.T) {
	response, err := private.GetServicesLimitsBuy(kamoney_sdk_dtos.GetServicesLimitsBuyRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetAffiliateInfo(t *testing.T) {
	response, err := private.GetAffiliateInfo(kamoney_sdk_dtos.GetAffiliateInfoRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetLevelInfo(t *testing.T) {
	response, err := private.GetLevelInfo(kamoney_sdk_dtos.GetLevelInfoRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetFee(t *testing.T) {
	response, err := private.GetFee(kamoney_sdk_dtos.GetFeeRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetReward(t *testing.T) {
	response, err := private.GetReward(kamoney_sdk_dtos.GetRewardRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListRecipients(t *testing.T) {
	response, err := private.ListRecipients(kamoney_sdk_dtos.ListRecipientsRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestCreateRecipients(t *testing.T) {
	response, err := private.CreateRecipients(kamoney_sdk_dtos.CreateRecipientsRequestParams{
		Type:          1,
		AccountType:   "CC",
		Bank:          1,
		Agency:        "9875",
		AccountNumber: "2222-3",
		Owner:         "Nome do Fulano",
		PersonalID:    "08830121622",
		Description:   "Test do SDK",
		Nonce:         fmt.Sprint(utility.GenNonce()),
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestDeleteRecipients(t *testing.T) {
	response, err := private.DeleteRecipients(kamoney_sdk_dtos.DeleteRecipientsRequestParams{}, "2911")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUpdateRecipients(t *testing.T) {
	response, err := private.ChangeRecipients(kamoney_sdk_dtos.ChangeRecipientsRequestParams{
		Type:          1,
		AccountType:   "CC",
		Bank:          1,
		Agency:        "0000",
		AccountNumber: "98765-3",
		Owner:         "João do Fulano",
		PersonalID:    "05650090602",
		Description:   "Test Update do SDK",
		Nonce:         fmt.Sprint(utility.GenNonce()),
	}, "2912")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetRecipientInfo(t *testing.T) {
	response, err := private.GetRecipientInfo(kamoney_sdk_dtos.GetRecipientInfoRequestParams{}, "2912")

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
