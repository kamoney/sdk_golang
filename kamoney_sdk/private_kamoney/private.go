package private_kamoney

import (
	"fmt"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

var (
	ENDPOINT_ACCOUNT_INFO            = "private/account"
	ENDPOINT_ACCOUNT_LOCALITY        = "private/account/locality"
	ENDPOINT_ACCOUNT_CONTACT         = "private/account/contact"
	ENDPOINT_ACCOUNT_HISTORY         = "private/account/history"
	ENDPOINT_ACCOUNT_NOTIFICATION    = "private/account/notification"
	ENDPOINT_ACCOUNT_NOTIFICATION_ID = func(id string) string {
		return fmt.Sprintf("private/account/notification/%v", id)
	}
	ENDPOINT_ACCOUNT_KYC                      = "private/account/kyc"
	ENDPOINT_ACCOUNT_SERVICES_LIMITS_ORDER    = "private/account/services/order"
	ENDPOINT_ACCOUNT_SERVICES_LIMITS_MERCHANT = "private/account/services/merchant"
	ENDPOINT_ACCOUNT_SERVICES_LIMITS_BUY      = "private/account/services/buy"
	ENDPOINT_ACCOUNT_AFFILIATES               = "private/account/affiliates"
	ENDPOINT_ACCOUNT_LEVEL                    = "private/account/limit"
	ENDPOINT_ACCOUNT_FEE                      = "private/account/fee"
	ENDPOINT_ACCOUNT_REWARD                   = "private/account/reward"
	ENDPOINT_ACCOUNT_RECIPIENTS               = "private/account/recipients"
	ENDPOINT_ACCOUNT_RECIPIENTS_ID            = func(id string) string {
		return fmt.Sprintf("private/account/recipients/%v", id)
	}

	ENDPOINT_SECURITY_ANTIPHISHING      = "private/security/antiphishing"
	ENDPOINT_SECURITY_ANTIPHISHING_VIEW = "private/security/antiphishing/view"
	ENDPOINT_SECURITY_TFS               = "private/security/tfs"
	ENDPOINT_SECURITY_EMAIL             = "private/security/email"
	ENDPOINT_SECURITY_PASSWORD          = "private/security/password"
	ENDPOINT_SECURITY_ACTION            = "private/security/action"
	ENDPOINT_SECURITY_CANCEL            = "private/security/cancel"
	ENDPOINT_SECURITY_API               = "private/security/api"
	ENDPOINT_SECURITY_API_DELETE        = func(id int64) string {
		return fmt.Sprintf("private/security/api/%v", id)
	}
	ENDPOINT_SECURITY_API_SECRET = func(id int64) string {
		return fmt.Sprintf("private/security/api/%v/secret", id)
	}

	ENDPOINT_WALLET        = "private/wallet"
	ENDPOINT_WITHDRAW      = "private/wallet/withdraw"
	ENDPOINT_WITHDRAW_INFO = func(id string) string {
		return fmt.Sprintf("private/wallet/withdraw/%v", id)
	}
	ENDPOINT_WITHDRAW_RECEIPT = func(id string) string {
		return fmt.Sprintf("private/wallet/withdraw/%v/receipt", id)
	}
	ENDPOINT_WITHDRAW_RECEIPT_DOWNLOAD = func(id, filename string) string {
		return fmt.Sprintf("private/wallet/withdraw/%v/receipt/%v", id, filename)
	}

	ENDPOINT_ORDER      = "private/order"
	ENDPOINT_ORDER_INFO = func(id string) string {
		return fmt.Sprintf("private/order/%v", id)
	}
	ENDPOINT_ORDER_RECEIPT          = "private/order/receipt"
	ENDPOINT_ORDER_RECEIPT_DOWNLOAD = func(filename string) string {
		return fmt.Sprintf("private/order/receipt/%v", filename)
	}

	ENDPOINT_BUY      = "private/buy"
	ENDPOINT_BUY_INFO = func(id string) string {
		return fmt.Sprintf("private/buy/%v", id)
	}
	ENDPOINT_BUY_QR_CODE = func(id string) string {
		return fmt.Sprintf("private/buy/%v/payment_method/reset", id)
	}
	ENDPOINT_BUY_PRIVATE = func(id string) string {
		return fmt.Sprintf("private/buy/%v/private", id)
	}

	ENDPOINT_MERCHANT      = "private/merchant"
	ENDPOINT_MERCHANT_INFO = func(id string) string {
		return fmt.Sprintf("private/merchant/%v", id)
	}

	ENDPOINT_PAYMENT_LINK        = "private/merchant/paymentlink"
	ENDPOINT_PAYMENT_LINK_DELETE = func(id string) string {
		return fmt.Sprintf("private/merchant/paymentlink/%v", id)
	}

// ENDPOINT_AUTH              = "public/auth"
//
//	ENDPOINT_UTILS_COUNTRY_STATE = func(id int64) string {
//		return fmt.Sprintf("public/country/%v/state", id)
//	}
)

type PrivateRequestsInterface interface {
	GetAccountInfo(kamoney_sdk_dtos.ChangeAccountInfoRequestParams) (kamoney_sdk_dtos.GetAccountInfoRequestResponse, error)
	ChangeAccountInfo(kamoney_sdk_dtos.ChangeAccountInfoRequestParams) (kamoney_sdk_dtos.ChangeAccountInfoRequestResponse, error)

	ListRecipients(kamoney_sdk_dtos.ListRecipientsRequestParams) (kamoney_sdk_dtos.ListRecipientsRequestResponse, error)
	CreateRecipients(kamoney_sdk_dtos.CreateRecipientsRequestParams) (kamoney_sdk_dtos.CreateRecipientsRequestResponse, error)
	DeleteRecipients(kamoney_sdk_dtos.DeleteRecipientsRequestParams, string) (kamoney_sdk_dtos.DeleteRecipientsRequestResponse, error)
	ChangeRecipients(kamoney_sdk_dtos.ChangeRecipientsRequestParams, string) (kamoney_sdk_dtos.ChangeRecipientsRequestResponse, error)
	GetRecipientInfo(kamoney_sdk_dtos.GetRecipientInfoRequestParams, string) (kamoney_sdk_dtos.GetRecipientInfoRequestResponse, error)

	GetAccountLocality(kamoney_sdk_dtos.GetAccountLocalityRequestParams) (kamoney_sdk_dtos.GetAccountLocalityRequestResponse, error)
	ChangeAccountLocality(kamoney_sdk_dtos.ChangeAccountLocalityRequestParams) (kamoney_sdk_dtos.ChangeAccountLocalityRequestResponse, error)

	ChangeAccountContact(kamoney_sdk_dtos.ChangeAccountContactRequestParams) (kamoney_sdk_dtos.ChangeAccountContactRequestResponse, error)
	GetAccountHistory(kamoney_sdk_dtos.GetAccountHistoryRequestParams) (kamoney_sdk_dtos.GetAccountHistoryRequestResponse, error)
	GetAccountNotification(kamoney_sdk_dtos.GetAccountNotificationRequestParams) (kamoney_sdk_dtos.GetAccountNotificationRequestResponse, error)
	UpdateAccountNotificationReadAll(kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestParams) (kamoney_sdk_dtos.UpdateAccountNotificationReadAllRequestResponse, error)
	UpdateAccountNotificationReadId(kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestParams, string) (kamoney_sdk_dtos.UpdateAccountNotificationReadIdRequestResponse, error)
	GetAccountKyc(kamoney_sdk_dtos.GetAccountKycRequestParams) (kamoney_sdk_dtos.GetAccountKycRequestResponse, error)

	GetServicesLimitsOrder(kamoney_sdk_dtos.GetServicesLimitsOrderRequestParams) (kamoney_sdk_dtos.GetServicesLimitsOrderRequestResponse, error)
	GetServicesLimitsMerchant(kamoney_sdk_dtos.GetServicesLimitsMerchantRequestParams) (kamoney_sdk_dtos.GetServicesLimitsMerchantRequestResponse, error)
	GetServicesLimitsBuy(kamoney_sdk_dtos.GetServicesLimitsBuyRequestParams) (kamoney_sdk_dtos.GetServicesLimitsBuyRequestResponse, error)

	GetAffiliateInfo(kamoney_sdk_dtos.GetAffiliateInfoRequestParams) (kamoney_sdk_dtos.GetAffiliateInfoRequestResponse, error)
	GetLevelInfo(kamoney_sdk_dtos.GetLevelInfoRequestParams) (kamoney_sdk_dtos.GetLevelInfoRequestResponse, error)
	GetFee(kamoney_sdk_dtos.GetFeeRequestParams) (kamoney_sdk_dtos.GetFeeRequestResponse, error)
	GetReward(kamoney_sdk_dtos.GetRewardRequestParams) (kamoney_sdk_dtos.GetRewardRequestResponse, error)

	ListWallet(kamoney_sdk_dtos.ListWalletRequestParams) (kamoney_sdk_dtos.ListWalletRequestResponse, error)
	GetWalletExtract(kamoney_sdk_dtos.GetWalletExtractRequestParams) (kamoney_sdk_dtos.GetWalletExtractRequestResponse, error)
	ChangeAntiPhishing(kamoney_sdk_dtos.ChangeAntiPhishingRequestParams) (kamoney_sdk_dtos.ChangeAntiPhishingRequestResponse, error)
	ViewAntiPhishing(kamoney_sdk_dtos.ViewAntiPhishingRequestParams) (kamoney_sdk_dtos.ViewAntiPhishingRequestResponse, error)

	// GetTFS(kamoney_sdk_dtos.GetTFSRequestParams) (kamoney_sdk_dtos.GetTFSRequestResponse, error)
	// CreateTfs(kamoney_sdk_dtos.CreateTfsRequestParams) (kamoney_sdk_dtos.CreateTfsRequestResponse, error)
	// ChangeEmail(kamoney_sdk_dtos.ChangeEmailRequestParams) (kamoney_sdk_dtos.ChangeEmailRequestResponse, error)
	// ChangePassword(kamoney_sdk_dtos.ChangePasswordRequestParams) (kamoney_sdk_dtos.ChangePasswordRequestResponse, error)
	// SecurityAction(kamoney_sdk_dtos.SecurityActionRequestParams) (kamoney_sdk_dtos.SecurityActionRequestResponse, error)
	// CancelAccount(kamoney_sdk_dtos.CancelAccountRequestParams) (kamoney_sdk_dtos.CancelAccountRequestResponse, error)

	// CreateAPI(kamoney_sdk_dtos.CreateAPIRequestParams) (kamoney_sdk_dtos.CreateAPIRequestResponse, error)
	// ListAPI(kamoney_sdk_dtos.ListAPIRequestParams) (kamoney_sdk_dtos.ListAPIRequestResponse, error)
	// DeleteAPI(kamoney_sdk_dtos.DeleteAPIRequestParams) (kamoney_sdk_dtos.DeleteAPIRequestResponse, error)
	// GetAPISecret(kamoney_sdk_dtos.GetAPISecretRequestParams) (kamoney_sdk_dtos.GetAPISecretRequestResponse, error)

	CreateWithdraw(kamoney_sdk_dtos.CreateWithdrawRequestParams) (kamoney_sdk_dtos.CreateWithdrawRequestResponse, error)
	ListWithdraw(kamoney_sdk_dtos.ListWithdrawRequestParams) (kamoney_sdk_dtos.ListWithdrawRequestResponse, error)
	GetWithdrawInfo(kamoney_sdk_dtos.GetWithdrawInfoRequestParams, string) (kamoney_sdk_dtos.GetWithdrawInfoRequestResponse, error)
	GetWithdrawReceipt(kamoney_sdk_dtos.GetWithdrawReceiptRequestParams, string) (kamoney_sdk_dtos.GetWithdrawReceiptRequestResponse, error)
	GetWithdrawReceiptDownload(kamoney_sdk_dtos.GetWithdrawReceiptDownloadRequestParams, string) (kamoney_sdk_dtos.GetWithdrawReceiptDownloadRequestResponse, error)

	CreateOrder(kamoney_sdk_dtos.CreateOrderRequestParams) (kamoney_sdk_dtos.CreateOrderRequestResponse, error)
	ListOrder(kamoney_sdk_dtos.ListOrderRequestParams) (kamoney_sdk_dtos.ListOrderRequestResponse, error)
	GetOrderInfo(kamoney_sdk_dtos.GetOrderInfoRequestParams) (kamoney_sdk_dtos.GetOrderInfoRequestResponse, error)
	ListOrderReceipt(kamoney_sdk_dtos.ListOrderReceiptRequestParams) (kamoney_sdk_dtos.ListOrderReceiptRequestResponse, error)
	GetOrderReceiptDownload(kamoney_sdk_dtos.GetOrderReceiptDownloadRequestParams) (kamoney_sdk_dtos.GetOrderReceiptDownloadRequestResponse, error)

	ListBuy(kamoney_sdk_dtos.ListBuyRequestParams) (kamoney_sdk_dtos.ListBuyRequestResponse, error)
	GetBuyInfo(kamoney_sdk_dtos.GetBuyInfoRequestParams, string) (kamoney_sdk_dtos.GetBuyInfoRequestResponse, error)
	// GetBuyNewQRCode(kamoney_sdk_dtos.GetBuyNewQRCodeRequestParams, string) (kamoney_sdk_dtos.GetBuyNewQRCodeRequestResponse, error)
	// GetBuyPrivateKey(kamoney_sdk_dtos.GetBuyPrivateKeyRequestParams, string) (kamoney_sdk_dtos.GetBuyPrivateKeyRequestResponse, error)
	CreateBuy(kamoney_sdk_dtos.CreateBuyRequestParams) (kamoney_sdk_dtos.CreateBuyRequestResponse, error)

	CreateMerchant(kamoney_sdk_dtos.CreateMerchantRequestParams) (kamoney_sdk_dtos.CreateMerchantRequestResponse, error)
	ListMerchant(kamoney_sdk_dtos.ListMerchantRequestParams) (kamoney_sdk_dtos.ListMerchantRequestResponse, error)
	GetMerchantInfo(kamoney_sdk_dtos.GetMerchantInfoRequestParams, string) (kamoney_sdk_dtos.GetMerchantInfoRequestResponse, error)

	CreatePaymentLink(kamoney_sdk_dtos.CreatePaymentLinkRequestParams) (kamoney_sdk_dtos.CreatePaymentLinkRequestResponse, error)
	ListPaymentLink(kamoney_sdk_dtos.ListPaymentLinkRequestParams) (kamoney_sdk_dtos.ListPaymentLinkRequestResponse, error)
	DeletePaymentLink(kamoney_sdk_dtos.DeletePaymentLinkRequestParams, int64) (kamoney_sdk_dtos.DeletePaymentLinkRequestResponse, error)
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
