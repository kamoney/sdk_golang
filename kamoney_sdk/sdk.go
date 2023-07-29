package kamoney_sdk

import (
	"github.com/kamoney/sdk_golang/kamoney_sdk/request"
	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

var (
	ENDPOINT_AUTH              = "public/auth"
	ENDPOINT_SERVICES_ORDER    = "public/services/order"
	ENDPOINT_SERVICES_MERCHANT = "public/services/merchant"
	ENDPOINT_SERVICES_BUY      = "public/services/buy"

	ENDPOINT_ACCOUNT_REGISTER         = "public/register"
	ENDPOINT_ACCOUNT_ACTIVE           = "public/active"
	ENDPOINT_ACCOUNT_RECOVERY         = "public/recovery"
	ENDPOINT_ACCOUNT_RECOVERY_CONFIRM = "public/recovery/confirm"
)

type sdk struct {
	email          string
	password       string
	publicKey      string
	secretKey      string
	requestHandler *request.RequestHandler
}

type SDK interface {
	// Auth(kamoney_sdk_dtos.AuthRequestParams) (kamoney_sdk_dtos.AuthRequestResponse, error)
	ServicesOrder() (kamoney_sdk_dtos.ServicesOrderRequestResponse, error)
	ServicesMerchant() (kamoney_sdk_dtos.ServicesMerchantRequestResponse, error)
	ServicesBuy() (kamoney_sdk_dtos.ServicesBuyRequestResponse, error)

	Register(kamoney_sdk_dtos.RegisterRequestParams) (kamoney_sdk_dtos.RegisterRequestResponse, error)
	Active(kamoney_sdk_dtos.ActiveRequestParams) (kamoney_sdk_dtos.ActiveRequestResponse, error)
	Recovery(kamoney_sdk_dtos.RecoveryRequestParams) (kamoney_sdk_dtos.RecoveryRequestResponse, error)
	RecoveryConfirm(kamoney_sdk_dtos.RecoveryConfirmRequestParams) (kamoney_sdk_dtos.RecoveryConfirmRequestResponse, error)
}

func NewSDKClient(email, password, publicKey, secretKey string) SDK {
	r := request.RequestHandler{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}

	return &sdk{
		email:          email,
		password:       password,
		requestHandler: &r,
	}
}
