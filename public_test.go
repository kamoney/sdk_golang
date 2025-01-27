package kamoney_sdk_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kamoney/sdk_golang/kamoney_sdk/private_kamoney"
	"github.com/kamoney/sdk_golang/kamoney_sdk/public_kamoney"
	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

var public public_kamoney.PublicRequestsInterface
var private private_kamoney.PrivateRequestsInterface

func TestMain(m *testing.M) {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatal(".env file is missing")
	}

	public = public_kamoney.NewPublicRequests(os.Getenv("PUBLIC_KEY"))
	private = private_kamoney.NewPrivateRequests(os.Getenv("EMAIL"), os.Getenv("PASS"), os.Getenv("PUBLIC_KEY"), os.Getenv("SECRET_KEY"))

	code := m.Run()

	fmt.Println("Limpando recursos globais...")

	os.Exit(code)
}

func TestServicesOrder(t *testing.T) {
	response, err := public.ServicesOrder()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestServicesMerchant(t *testing.T) {
	response, err := public.ServicesMerchant()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestServicesBuy(t *testing.T) {
	response, err := public.ServicesBuy()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestAccountRegister(t *testing.T) {
	response, err := public.AccountRegister(kamoney_sdk_dtos.AccountRegisterRequestParams{
		Email:         "immortal.g.tv@gmail.com",
		AffiliateCode: "QGHZ0RYU5R",
		Terms:         true,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestAccountActive(t *testing.T) {
	response, err := public.AccountActive(kamoney_sdk_dtos.AccountActiveRequestParams{
		Email:    "immortal.g.tv@gmail.com",
		Code:     324340,
		Password: "m1nh@s3nh@n0v@",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestAccountRecovery(t *testing.T) {
	response, err := public.AccountRecovery(kamoney_sdk_dtos.AccountRecoveryRequestParams{
		Email: "immortal.g.tv@gmail.com",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestAccountRecoveryConfirm(t *testing.T) {
	response, err := public.AccountRecoveryConfirm(kamoney_sdk_dtos.AccountRecoveryConfirmRequestParams{
		Email:    "immortal.g.tv@gmail.com",
		Code:     188636,
		Password: "7023346aA@",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsBanks(t *testing.T) {
	response, err := public.UtilsBanks()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsNotification(t *testing.T) {
	response, err := public.UtilsNotification()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsCountry(t *testing.T) {
	response, err := public.UtilsCountry()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsState(t *testing.T) {
	response, err := public.UtilsState(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsCity(t *testing.T) {
	response, err := public.UtilsCity(1, 19)
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsCurrency(t *testing.T) {
	response, err := public.UtilsCurrency()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsCurrencyNetwork(t *testing.T) {
	response, err := public.UtilsCurrencyNetwork("")
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsFaq(t *testing.T) {
	_, err := public.UtilsFaq()
	if err != nil {
		t.Error(err)
	}

	// t.Log(response)
}

func TestUtilsProduct(t *testing.T) {
	response, err := public.UtilsProduct()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsContact(t *testing.T) {
	response, err := public.UtilsContact(kamoney_sdk_dtos.UtilsContactRequestParams{
		Name:    "Igor",
		Email:   "immortal.g.tv@gmail.com",
		Subject: "Test API",
		Message: "Olá, só passando pra dizer Oi :D",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestUtilsPixType(t *testing.T) {
	response, err := public.UtilsPixType()
	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
