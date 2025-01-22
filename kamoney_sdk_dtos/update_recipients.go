package kamoney_sdk_dtos

type UpdateRecipientsRequestParams struct {
	Nonce         string `json:"nonce"`
	ID            int64  `json:"id"`
	Type          int64  `json:"type"`
	AccountType   string `json:"account_type"`
	Bank          int64  `json:"bank"`
	Agency        string `json:"agency"`
	AccountNumber string `json:"account_number"`
	Owner         string `json:"owner"`
	PersonalID    string `json:"personal_id"`
	Description   string `json:"description"`
}

type UpdateRecipientsRequestResponse struct {
	Common
}
