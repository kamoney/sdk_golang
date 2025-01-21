package private_kamoney

import (
	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

var (
	ENDPOINT_ACCOUNT_INFO               = "private/account"
	ENDPOINT_ACCOUNT_LOCALITY           = "private/account/locality"
	ENDPOINT_ACCOUNT_CONTACT            = "private/account/contact"
	ENDPOINT_ACCOUNT_HISTORY            = "private/account/history"
	ENDPOINT_ACCOUNT_NOTIFICATION       = "private/account/notification"
	ENDPOINT_ACCOUNT_KYC                = "private/account/kyc"
	ENDPOINT_SERVICES_LIMITS_ORDER      = "private/account/services/order"
	ENDPOINT_SERVICES_LIMITS_MERCHANT   = "private/account/services/merchant"
	ENDPOINT_SERVICES_LIMITS_BUY        = "private/account/services/buy"
	ENDPOINT_SERVICES_AFFILIATES        = "private/account/affiliates"
	ENDPOINT_SERVICES_LEVEL             = "private/account/limit"
	ENDPOINT_SERVICES_FEE               = "private/account/fee"
	ENDPOINT_SERVICES_REWARD            = "private/account/reward"
	ENDPOINT_SERVICES_WALLET            = "private/wallet"
	ENDPOINT_SERVICES_ANTIPHISHING      = "private/security/antiphishing"
	ENDPOINT_SERVICES_ANTIPHISHING_VIEW = "private/security/antiphishing/view"
	ENDPOINT_SERVICES_TFS               = "private/security/tfs"
	ENDPOINT_SERVICES_EMAIL             = "private/security/email"
	ENDPOINT_SERVICES_PASSWORD          = "private/security/password"
	ENDPOINT_SERVICES_ACTION            = "private/security/action"
	ENDPOINT_SERVICES_CANCEL            = "private/security/cancel"
	ENDPOINT_SERVICES_API               = "private/security/api"

// ENDPOINT_AUTH              = "public/auth"
//
//	ENDPOINT_UTILS_COUNTRY_STATE = func(id int64) string {
//		return fmt.Sprintf("public/country/%v/state", id)
//	}
)

type PrivateRequestsInterface interface {
	AccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams) (kamoney_sdk_dtos.AccountInfoRequestResponse, error)
	GetAccountInfo(kamoney_sdk_dtos.AccountInfoRequestParams) (kamoney_sdk_dtos.GetAccountInfoRequestResponse, error)

	GetAccountLocality(kamoney_sdk_dtos.GetAccountLocalityRequestParams) (kamoney_sdk_dtos.GetAccountLocalityRequestResponse, error)
	AccountLocality(kamoney_sdk_dtos.AccountLocalityRequestParams) (kamoney_sdk_dtos.AccountLocalityRequestResponse, error)

	AccountContact(kamoney_sdk_dtos.AccountContactRequestParams) (kamoney_sdk_dtos.AccountContactRequestResponse, error)
	GetAccountHistory(kamoney_sdk_dtos.GetAccountHistoryRequestParams) (kamoney_sdk_dtos.GetAccountHistoryRequestResponse, error)
	GetAccountNotification(kamoney_sdk_dtos.GetAccountNotificationRequestParams) (kamoney_sdk_dtos.GetAccountNotificationRequestResponse, error)
	UpdateAccountNotificationReadAll(kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestParams) (kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestResponse, error)
	UpdateAccountNotificationReadId(kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestParams) (kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestResponse, error)
	GetAccountKyc(kamoney_sdk_dtos.GetAccountKycRequestParams) (kamoney_sdk_dtos.GetAccountKycRequestResponse, error)

	GetServicesLimitsOrder(kamoney_sdk_dtos.GetServicesLimitsOrderRequestParams) (kamoney_sdk_dtos.GetServicesLimitsOrderRequestResponse, error)
	GetServicesLimitsMerchant(kamoney_sdk_dtos.GetServicesLimitsMerchantRequestParams) (kamoney_sdk_dtos.GetServicesLimitsMerchantRequestResponse, error)
	GetServicesLimitsBuy(kamoney_sdk_dtos.GetServicesLimitsBuyRequestParams) (kamoney_sdk_dtos.GetServicesLimitsBuyRequestResponse, error)

	GetAffiliateInfo(kamoney_sdk_dtos.GetAffiliateInfoRequestParams) (kamoney_sdk_dtos.GetAffiliateInfoRequestResponse, error)
	GetLevelInfo(kamoney_sdk_dtos.GetLevelInfoRequestParams) (kamoney_sdk_dtos.GetLevelInfoRequestResponse, error)
	GetFee(kamoney_sdk_dtos.GetFeeRequestParams) (kamoney_sdk_dtos.GetFeeRequestResponse, error)
	GetReward(kamoney_sdk_dtos.GetRewardRequestParams) (kamoney_sdk_dtos.GetRewardRequestResponse, error)

	ListWallet(kamoney_sdk_dtos.ListWalletRequestParams) (kamoney_sdk_dtos.ListWalletRequestResponse, error)
	WalletExtract(kamoney_sdk_dtos.WalletExtractRequestParams) (kamoney_sdk_dtos.WalletExtractRequestResponse, error)
	AntiPhishing(kamoney_sdk_dtos.AntiPhishingRequestParams) (kamoney_sdk_dtos.AntiPhishingRequestResponse, error)
	ViewAntiPhishing(kamoney_sdk_dtos.ViewAntiPhishingRequestParams) (kamoney_sdk_dtos.ViewAntiPhishingRequestResponse, error)
	ViewTfs(kamoney_sdk_dtos.ViewTfsRequestParams) (kamoney_sdk_dtos.ViewTfsRequestResponse, error)
	CreateTfs(kamoney_sdk_dtos.CreateTfsRequestParams) (kamoney_sdk_dtos.CreateTfsRequestResponse, error)
	ChangeEmail(kamoney_sdk_dtos.ChangeEmailRequestParams) (kamoney_sdk_dtos.ChangeEmailRequestResponse, error)
	ChangePassword(kamoney_sdk_dtos.ChangePasswordRequestParams) (kamoney_sdk_dtos.ChangePasswordRequestResponse, error)
	SecurityAction(kamoney_sdk_dtos.SecurityActionRequestParams) (kamoney_sdk_dtos.SecurityActionRequestResponse, error)
	CancelAccount(kamoney_sdk_dtos.CancelAccountRequestParams) (kamoney_sdk_dtos.CancelAccountRequestResponse, error)

	CreateAPI(kamoney_sdk_dtos.CreateAPIRequestParams) (kamoney_sdk_dtos.CreateAPIRequestResponse, error)
	ListAPIs(kamoney_sdk_dtos.ListAPIsRequestParams) (kamoney_sdk_dtos.ListAPIsRequestResponse, error)
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
