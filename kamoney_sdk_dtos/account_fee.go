package kamoney_sdk_dtos

type GetFeeRequestParams struct {
}

type GetFeeRequestResponse struct {
	Common
	Data struct {
		Order struct {
			BTC struct {
				Fee              float64 `json:"fee"`
				Promotion        bool    `json:"promotion"`
				NextLevel        bool    `json:"next_level"`
				NextLevelMissing float64 `json:"next_level_missing"`
				NextLevelFee     float64 `json:"next_level_fee"`
			} `json:"BTC"`
			// Adicione outras moedas aqui conforme necessário
		} `json:"order"`
		Merchant struct {
			BTC struct {
				Fee float64 `json:"fee"`
			} `json:"BTC"`
			// Adicione outras moedas aqui conforme necessário
		} `json:"merchant"`
		Buy struct {
			BTC struct {
				Fee float64 `json:"fee"`
			} `json:"BTC"`
			// Adicione outras moedas aqui conforme necessário
		} `json:"buy"`
		Deposit  []interface{} `json:"deposit"` // Se "deposit" for uma lista vazia, podemos usar interface{}
		Withdraw struct {
			BRL struct {
				Variable float64 `json:"variable"`
				Fixed    float64 `json:"fixed"`
			} `json:"BRL"`
		} `json:"withdraw"`
	} `json:"data"`
}
