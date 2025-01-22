package kamoney_sdk_dtos

type CreateRecipientsRequestParams struct {
	Nonce         string `json:"nonce"`
	Type          int64  `json:"type"`
	AccountType   string `json:"account_type"`
	Bank          int64  `json:"bank"`
	Agency        string `json:"agency"`
	AccountNumber string `json:"account_number"`
	Owner         string `json:"owner"`
	PersonalID    string `json:"personal_id"`
	Description   string `json:"description"`
}

type CreateRecipientsRequestResponse struct {
	Common
}
