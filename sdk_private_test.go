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

// Only through portal access
func TestDeleteAPI(t *testing.T) {
	response, err := private.DeleteAPI(kamoney_sdk_dtos.DeleteAPIRequestParams{
		ID:       1,
		Password: "",
		Tfa:      1,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// Only through portal access
func TestGetAPISecret(t *testing.T) {
	response, err := private.GetAPISecret(kamoney_sdk_dtos.GetAPISecretRequestParams{
		ID:       1,
		Password: "",
		Tfa:      1,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestListRecipients(t *testing.T) {
	response, err := private.ListRecipients(kamoney_sdk_dtos.ListRecipientsRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
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
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestDeleteRecipients(t *testing.T) {
	response, err := private.DeleteRecipients(kamoney_sdk_dtos.DeleteRecipientsRequestParams{
		ID: 2908,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

// endpoint still Exists ? ERR 400
func TestUpdateRecipients(t *testing.T) {
	response, err := private.UpdateRecipients(kamoney_sdk_dtos.UpdateRecipientsRequestParams{
		ID:            2908,
		Type:          1,
		AccountType:   "CC",
		Bank:          1,
		Agency:        "0000",
		AccountNumber: "98765-3",
		Owner:         "João do Fulano",
		PersonalID:    "05650090602",
		Description:   "Test Update do SDK",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetRecipientInfo(t *testing.T) {
	response, err := private.GetRecipientInfo(kamoney_sdk_dtos.GetRecipientInfoRequestParams{
		AccountType:    1,
		Bank:           1,
		Agency:         "0000",
		AccountNumber:  "98765-0",
		Owner:          "Claudeci Goularte",
		PersonalID:     "05650090602",
		Description:    "Novo cadastro.3 16-05-2022",
		Main:           0,
		DDD:            "31",
		MobileOperator: "Claro",
		PhoneNumber:    "982568095",
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

func TestListOrder(t *testing.T) {
	response, err := private.ListOrder(kamoney_sdk_dtos.ListOrderRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
