package kamoney_sdk_dtos

type GetRecipientInfoRequestParams struct{}

type GetRecipientInfoRequestResponse struct {
	Common
	Data struct {
		ID                     int64  `json:"id"`
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
	} `json:"data"`
}
