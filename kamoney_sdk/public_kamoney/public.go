package public_kamoney

import (
	"fmt"

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
	ENDPOINT_UTILS_BANK               = "public/bank"
	ENDPOINT_UTILS_NOTIFICATION       = "public/notification"
	ENDPOINT_UTILS_COUNTRY            = "public/country"
	ENDPOINT_UTILS_COUNTRY_STATE      = func(id int64) string {
		return fmt.Sprintf("public/country/%v/state", id)
	}
	ENDPOINT_UTILS_COUNTRY_CITY = func(country_id, state_id int64) string {
		return fmt.Sprintf("public/country/%v/state/%v/city", country_id, state_id)
	}
	ENDPOINT_UTILS_CURRENCY         = "public/currency"
	ENDPOINT_UTILS_CURRENCY_NETWORK = func(asset string) string {
		return fmt.Sprintf("public/currency/%v", asset)
	}
	ENDPOINT_UTILS_FAQ      = "public/faq"
	ENDPOINT_UTILS_PRODUCT  = "public/product"
	ENDPOINT_UTILS_CONTACT  = "public/contact"
	ENDPOINT_UTILS_PIX_TYPE = "public/pixtype"
)

type PublicRequestsInterface interface {
	// PUBLIC
	// Auth(kamoney_sdk_dtos.AuthRequestParams) (kamoney_sdk_dtos.AuthRequestResponse, error)
	ServicesOrder() (kamoney_sdk_dtos.ServicesOrderRequestResponse, error)
	ServicesMerchant() (kamoney_sdk_dtos.ServicesMerchantRequestResponse, error)
	ServicesBuy() (kamoney_sdk_dtos.ServicesBuyRequestResponse, error)

	AccountRegister(kamoney_sdk_dtos.AccountRegisterRequestParams) (kamoney_sdk_dtos.AccountRegisterRequestResponse, error)
	AccountActive(kamoney_sdk_dtos.AccountActiveRequestParams) (kamoney_sdk_dtos.AccountActiveRequestResponse, error)
	AccountRecovery(kamoney_sdk_dtos.AccountRecoveryRequestParams) (kamoney_sdk_dtos.AccountRecoveryRequestResponse, error)
	AccountRecoveryConfirm(kamoney_sdk_dtos.AccountRecoveryConfirmRequestParams) (kamoney_sdk_dtos.AccountRecoveryConfirmRequestResponse, error)

	UtilsBanks() (kamoney_sdk_dtos.UtilsBankRequestResponse, error)
	UtilsNotification() (kamoney_sdk_dtos.UtilsNotificationRequestResponse, error)
	UtilsCountry() (kamoney_sdk_dtos.UtilsCountryRequestResponse, error)
	UtilsState(int64) (kamoney_sdk_dtos.UtilsStateRequestResponse, error)
	UtilsCity(int64, int64) (kamoney_sdk_dtos.UtilsCityRequestResponse, error)
	UtilsCurrency() (kamoney_sdk_dtos.UtilsCurrencyRequestResponse, error)
	UtilsCurrencyNetwork(string) (kamoney_sdk_dtos.UtilsCurrencyNetworkRequestResponse, error)
	UtilsFaq() (kamoney_sdk_dtos.UtilsFaqRequestResponse, error)
	UtilsProduct() (kamoney_sdk_dtos.UtilsProductRequestResponse, error)
	UtilsContact(kamoney_sdk_dtos.UtilsContactRequestParams) (kamoney_sdk_dtos.UtilsContactRequestResponse, error)
	UtilsPixType() (kamoney_sdk_dtos.UtilsPixTypeRequestResponse, error)
}

type publicRequests struct {
	PublicKey string
	r         *RequestHandler
}

func NewPublicRequests(publicKey string) PublicRequestsInterface {
	r := RequestHandler{
		PublicKey: publicKey,
	}

	return &publicRequests{
		r: &r,
	}
}
