package kamoney_sdk_dtos

type GetWithdrawInfoRequestParams struct {
	Nonce string `json:"nonce"`
	ID    string `json:"id"`
}

type GetWithdrawInfoRequestResponse struct {
	Common
	Data struct {
		ID       string  `json:"id"`
		Created  string  `json:"created"`
		Symbol   string  `json:"symbol"`
		Amount   float64 `json:"amount"`
		Fee      float64 `json:"fee"`
		Total    float64 `json:"total"`
		Modality int64   `json:"modality"`
		BankName string  `json:"bank_name"`
		Owner    struct {
			Name       string `json:"name"`
			PersonalID string `json:"personalid"`
		} `json:"owner"`
		Account struct {
			Number string `json:"number"`
			Type   string `json:"type"`
			Agency string `json:"agency"`
		} `json:"account"`
		Status struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"status"`
	} `json:"data"`
}
