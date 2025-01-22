package kamoney_sdk_dtos

type AntiPhishingRequestParams struct {
	Password string `json:"password"`
	Phrase   string `json:"phrase"`
	Tfa      int64  `json:"tfa"` //in case is active
	Nonce    string `json:"nonce"`
}

type AntiPhishingRequestResponse struct {
	Common
	Data []struct {
		Created         string  `json:"created"`
		History         string  `json:"history"`
		BalancePrevious float64 `json:"balance_previous"`
		Type            int64   `json:"type"`
		Amount          float64 `json:"amount"`
		Balance         float64 `json:"balance"`
		Status          struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"status"`
		ExpiresIn       string  `json:"expires_in"`
		AmountUsed      float64 `json:"amount_used"`
		AmountRemaining float64 `json:"amount_remaining"`
		TypeRegister    struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"type_register"`
	} `json:"data"`
}
