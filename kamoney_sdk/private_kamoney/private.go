package private_kamoney

import "github.com/kamoney/sdk_golang/kamoney_sdk_dtos"

var (
	ENDPOINT_ACCOUNT_INFO     = "private/account"
	ENDPOINT_ACCOUNT_LOCALITY = "account/locality"

// ENDPOINT_AUTH              = "public/auth"
//
//	ENDPOINT_UTILS_COUNTRY_STATE = func(id int64) string {
//		return fmt.Sprintf("public/country/%v/state", id)
//	}
)

type PrivateRequestsInterface interface {
	AccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams) (kamoney_sdk_dtos.AccountInfoRequestResponse, error)
	AccountLocality(kamoney_sdk_dtos.AccountLocalityRequestParams) (kamoney_sdk_dtos.AccountLocalityRequestResponse, error)
}

type privateRequests struct {
	Email     string
	Password  string
	PublicKey string
	SecretKey string
	r         *RequestHandler
}

func NewPrivateRequests(email, password, publicKey, secretKey string) PrivateRequestsInterface {
	r := RequestHandler{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}

	return &privateRequests{
		Email:     email,
		Password:  password,
		PublicKey: publicKey,
		SecretKey: secretKey,
		r:         &r,
	}
}
