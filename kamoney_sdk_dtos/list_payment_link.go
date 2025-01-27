package kamoney_sdk_dtos

type ListPaymentLinkRequestParams struct {
	Nonce  string `json:"nonce"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Search string `json:"search"`
	Status string `json:"status"`
}

type ListPaymentLinkRequestResponse struct {
	Common
	Data []struct {
		ID      int     `json:"id"`
		Created string  `json:"created"`
		Label   string  `json:"label"`
		Amount  float64 `json:"amount"`
		Link    string  `json:"link"`
	} `json:"data"`
}
