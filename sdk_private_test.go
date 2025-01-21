package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func TestGetAccountInfo(t *testing.T) {
	response, err := private.GetAccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountInfo(t *testing.T) {
	response, err := private.AccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams{
		Name:        "Igor Araújo da Silva",
		PersonalID:  "08830121622",
		DateOfBirth: "1995-02-05",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetAccountLocality(t *testing.T) {
	response, err := private.GetAccountLocality(kamoney_sdk_dtos.GetAccountLocalityRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

/* REDO: Invalid State param */
func TestAccountLocality(t *testing.T) {
	response, err := private.AccountLocality(kamoney_sdk_dtos.AccountLocalityRequestParams{
		Zipcode:      "35430232",
		Street:       "Av. Dom Bosco",
		Number:       490,
		Complement:   "201",
		Neighborhood: "Palmeiras",
		City:         79,
		State:        1,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountContact(t *testing.T) {
	response, err := private.AccountContact(kamoney_sdk_dtos.AccountContactRequestParams{
		Whatsapp: "986621962",
		Telegram: "@immortal",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetAccountHistory(t *testing.T) {
	response, err := private.GetAccountHistory(kamoney_sdk_dtos.GetAccountHistoryRequestParams{
		Page: 1,
		// Date: ,

	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
func TestGetAccountNotification(t *testing.T) {
	response, err := private.GetAccountNotification(kamoney_sdk_dtos.GetAccountNotificationRequestParams{
		Page: 1,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUpdateAccountNotificationReadAll(t *testing.T) {
	response, err := private.UpdateAccountNotificationReadAll(kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUpdateAccountNotificationReadId(t *testing.T) {
	response, err := private.UpdateAccountNotificationReadId(kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestParams{
		ID: "2936",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetAccountKyc(t *testing.T) {
	response, err := private.GetAccountKyc(kamoney_sdk_dtos.GetAccountKycRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetServicesLimitsOrder(t *testing.T) {
	response, err := private.GetServicesLimitsOrder(kamoney_sdk_dtos.GetServicesLimitsOrderRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetServicesLimitsMerchant(t *testing.T) {
	response, err := private.GetServicesLimitsMerchant(kamoney_sdk_dtos.GetServicesLimitsMerchantRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetServicesLimitsBuy(t *testing.T) {
	response, err := private.GetServicesLimitsBuy(kamoney_sdk_dtos.GetServicesLimitsBuyRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetAffiliateInfo(t *testing.T) {
	response, err := private.GetAffiliateInfo(kamoney_sdk_dtos.GetAffiliateInfoRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetLevelInfo(t *testing.T) {
	response, err := private.GetLevelInfo(kamoney_sdk_dtos.GetLevelInfoRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetFee(t *testing.T) {
	response, err := private.GetFee(kamoney_sdk_dtos.GetFeeRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetReward(t *testing.T) {
	response, err := private.GetReward(kamoney_sdk_dtos.GetRewardRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

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

func TestSecurityAntiPhishing(t *testing.T) {
	response, err := private.AntiPhishing(kamoney_sdk_dtos.AntiPhishingRequestParams{
		Password: "7023346a",
		Phrase:   "Abacaxi é fruta ?",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestSecurityViewAntiPhishing(t *testing.T) {
	response, err := private.ViewAntiPhishing(kamoney_sdk_dtos.ViewAntiPhishingRequestParams{
		Password: "7023346a",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Only through portal access
func TestSecurityViewTfs(t *testing.T) {
	response, err := private.ViewTfs(kamoney_sdk_dtos.ViewTfsRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Only through portal access
func TestSecurityCreateTfs(t *testing.T) {
	response, err := private.CreateTfs(kamoney_sdk_dtos.CreateTfsRequestParams{
		Password: "7023346a",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestSecurityEmail(t *testing.T) {
	response, err := private.ChangeEmail(kamoney_sdk_dtos.ChangeEmailRequestParams{
		Password: "7023346a",
		Email:    "immortal.g.tv@gmail.com",
		Tfa:      "123456",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestSecurityChangePassword(t *testing.T) {
	response, err := private.ChangePassword(kamoney_sdk_dtos.ChangePasswordRequestParams{
		Password:        "7023346a",
		PasswordNew:     "7023346aA@",
		PasswordConfirm: "7023346aA@",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestSecurityAction(t *testing.T) {
	response, err := private.SecurityAction(kamoney_sdk_dtos.SecurityActionRequestParams{
		Code: "730542",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Not tested
func TestCancelAccount(t *testing.T) {
	response, err := private.CancelAccount(kamoney_sdk_dtos.CancelAccountRequestParams{
		Password: "",
		Terms:    1,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Only through portal access
func TestCreateAPI(t *testing.T) {
	response, err := private.CreateAPI(kamoney_sdk_dtos.CreateAPIRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Only through portal access
func TestListAPIs(t *testing.T) {
	response, err := private.ListAPIs(kamoney_sdk_dtos.ListAPIsRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
