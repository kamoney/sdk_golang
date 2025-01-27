package kamoney_sdk_dtos

type DeletePaymentLinkRequestParams struct {
	Nonce  string `json:"nonce"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Search string `json:"search"`
	Status string `json:"status"`
}

type DeletePaymentLinkRequestResponse struct {
	Common
}
