package kamoney_sdk_dtos

type DeleteRecipientsRequestParams struct {
	Nonce string `json:"nonce"`
	ID    int64  `json:"id"`
}

type DeleteRecipientsRequestResponse struct {
	Common
}
