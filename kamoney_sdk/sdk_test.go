package kamoney_sdk_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kamoney/sdk_golang/kamoney_sdk"
	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

var sdk kamoney_sdk.SDK

func TestMain(m *testing.M) {
	err := godotenv.Load(`C:\Users\immor\OneDrive\Documentos\traderBot\.env`)

	if err != nil {
		log.Fatal(".env file is missing")
	}

	sdk = kamoney_sdk.NewSDKClient(os.Getenv("EMAIL"), os.Getenv("PASS"), os.Getenv("PUBLIC"), os.Getenv("SECRET"))

	code := m.Run()

	fmt.Println("Limpando recursos globais...")

	os.Exit(code)
}

func TestServicesOrder(t *testing.T) {
	response, err := sdk.ServicesOrder()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestServicesMerchant(t *testing.T) {
	response, err := sdk.ServicesMerchant()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestServicesBuy(t *testing.T) {
	response, err := sdk.ServicesBuy()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountRegister(t *testing.T) {
	response, err := sdk.AccountRegister(kamoney_sdk_dtos.AccountRegisterRequestParams{
		Email:         "immortal.g.tv@gmail.com",
		AffiliateCode: "QGHZ0RYU5R",
		Terms:         true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountActive(t *testing.T) {
	response, err := sdk.AccountActive(kamoney_sdk_dtos.AccountActiveRequestParams{
		Email:    "immortal.g.tv@gmail.com",
		Code:     324340,
		Password: "m1nh@s3nh@n0v@",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountRecovery(t *testing.T) {
	response, err := sdk.AccountRecovery(kamoney_sdk_dtos.AccountRecoveryRequestParams{
		Email: "immortal.g.tv@gmail.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestAccountRecoveryConfirm(t *testing.T) {
	response, err := sdk.AccountRecoveryConfirm(kamoney_sdk_dtos.AccountRecoveryConfirmRequestParams{
		Email:    "immortal.g.tv@gmail.com",
		Code:     324340,
		Password: "m1nh@s3nh@n0v@",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsBanks(t *testing.T) {
	response, err := sdk.UtilsBanks()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsNotification(t *testing.T) {
	response, err := sdk.UtilsNotification()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsCountry(t *testing.T) {
	response, err := sdk.UtilsCountry()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsState(t *testing.T) {
	response, err := sdk.UtilsState(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsCity(t *testing.T) {
	response, err := sdk.UtilsCity(1, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsCurrency(t *testing.T) {
	response, err := sdk.UtilsCurrency()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsCurrencyNetwork(t *testing.T) {
	response, err := sdk.UtilsCurrencyNetwork("")
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsFaq(t *testing.T) {
	_, err := sdk.UtilsFaq()
	if err != nil {
		panic(err)
	}

	// fmt.Println(response)
}

func TestUtilsProduct(t *testing.T) {
	response, err := sdk.UtilsProduct()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsContact(t *testing.T) {
	response, err := sdk.UtilsContact(kamoney_sdk_dtos.UtilsContactRequestParams{
		Name:    "Igor",
		Email:   "immortal.g.tv@gmail.com",
		Subject: "Test API",
		Message: "Olá, só passando pra dizer Oi :D",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestUtilsPixType(t *testing.T) {
	response, err := sdk.UtilsPixType()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
