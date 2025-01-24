package kamoney_sdk_dtos

type ListWalletRequestParams struct {
}

type ListWalletRequestResponse struct {
	Common
	Data []struct {
		Asset           string  `json:"asset"`
		Symbol          string  `json:"symbol"`
		Name            string  `json:"name"`
		In              float64 `json:"in"`
		Out             float64 `json:"out"`
		Blocked         float64 `json:"blocked"`
		Available       float64 `json:"available"`
		WithdrawEnabled bool    `json:"withdraw_enabled"`
	} `json:"data"`
}
