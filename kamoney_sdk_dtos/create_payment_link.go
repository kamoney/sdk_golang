package kamoney_sdk_dtos

type CreatePaymentLinkRequestParams struct {
	Nonce  string  `json:"nonce"`
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}

type CreatePaymentLinkRequestResponse struct {
	Common
	Data struct {
		ID      int     `json:"id"`
		Created string  `json:"created"`
		Label   string  `json:"label"`
		Amount  float64 `json:"amount"`
		Link    string  `json:"link"`
	} `json:"data"`
}
