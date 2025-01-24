package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

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
