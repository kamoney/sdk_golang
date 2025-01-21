package kamoney_sdk_dtos

type GetServicesLimitsBuyRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetServicesLimitsBuyRequestResponse struct {
	Common
	Data struct {
		Code         string   `json:"code"`
		Name         string   `json:"name"`
		Maintenance  bool     `json:"maintenance"`
		LimitMin     int      `json:"limit_min"`
		LimitMax     int      `json:"limit_max"`
		FeeThreshold int      `json:"fee_threshold"`
		FeeAmount    int      `json:"fee_amount"`
		Automatic    bool     `json:"automatic"`
		AutomaticMax int      `json:"automatic_max"`
		Info         []string `json:"info"`
		Limits       []struct {
			Owner int `json:"owner"`
			Other int `json:"other"`
		} `json:"limits"`
	} `json:"data"`
}
