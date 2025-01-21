package kamoney_sdk_dtos

type GetLevelInfoRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetLevelInfoRequestResponse struct {
	Common
	Data struct {
		UserID int `json:"user_id"`
		Level  struct {
			ID                 int    `json:"id"`
			Name               string `json:"name"`
			VerificationUpdate int    `json:"verification_update"`
		} `json:"level"`
		Verify struct {
			Enabled bool `json:"enabled"`
			Status  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"status"`
			Message string `json:"message"`
		} `json:"verify"`
		Limits []struct {
			Asset  string `json:"asset"`
			Symbol string `json:"symbol"`
			Name   string `json:"name"`
			Limit  struct {
				Sell struct {
					Total     float64 `json:"total"`
					Used      float64 `json:"used"`
					Remaining float64 `json:"remaining"`
				} `json:"sell"`
				Buy struct {
					Total     float64 `json:"total"`
					Used      float64 `json:"used"`
					Remaining float64 `json:"remaining"`
				} `json:"buy"`
			} `json:"limit"`
			Deposit struct {
				Enabled bool `json:"enabled"`
				Max     any  `json:"max"` // pode ser uma string se for um valor como "-"
			} `json:"deposit"`
			Withdraw struct {
				Enabled   bool    `json:"enabled"`
				Max       any     `json:"max"` // pode ser uma string
				Available float64 `json:"available"`
			} `json:"withdraw"`
		} `json:"limits"`
	} `json:"data"`
}
