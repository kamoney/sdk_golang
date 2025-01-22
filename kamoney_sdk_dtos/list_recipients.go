package kamoney_sdk_dtos

type ListRecipientsRequestParams struct {
	Nonce string `json:"nonce"`
}

type ListRecipientsRequestResponse struct {
	Common
	Data struct {
		BankAccount []struct {
			ID                     int    `json:"id"`
			Main                   bool   `json:"main"`
			AccountType            string `json:"account_type"`
			AccountTypeDescription string `json:"account_type_description"`
			Bank                   int64  `json:"bank"`
			BankName               string `json:"bank_name"`
			BankCode               string `json:"bank_code"`
			Agency                 string `json:"agency"`
			AccountNumber          string `json:"account_number"`
			PersonalID             string `json:"personal_id"`
			Owner                  string `json:"owner"`
			Description            string `json:"description"`
		} `json:"bank_account"`
		Phone []interface{} `json:"phone"`
	}
}
