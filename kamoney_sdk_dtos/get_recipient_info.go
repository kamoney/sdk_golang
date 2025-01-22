package kamoney_sdk_dtos

type GetRecipientInfoRequestParams struct {
	Nonce          string `json:"nonce"`
	ID             int64  `json:"id"`
	AccountType    int64  `json:"account_type"`
	Bank           int64  `json:"bank"`
	Agency         string `json:"agency"`
	AccountNumber  string `json:"account_number"`
	Owner          string `json:"owner"`
	PersonalID     string `json:"personalid"`
	Description    string `json:"description"`
	Main           int64  `json:"main"`
	DDD            string `json:"ddd"`
	MobileOperator string `json:"mobile_operator"`
	PhoneNumber    string `json:"phone_number"`
}

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
