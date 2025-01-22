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
		LimitMin     int64    `json:"limit_min"`
		LimitMax     int64    `json:"limit_max"`
		FeeThreshold int64    `json:"fee_threshold"`
		FeeAmount    int64    `json:"fee_amount"`
		Automatic    bool     `json:"automatic"`
		AutomaticMax int64    `json:"automatic_max"`
		Info         []string `json:"info"`
		Limits       []struct {
			Owner int64 `json:"owner"`
			Other int64 `json:"other"`
		} `json:"limits"`
	} `json:"data"`
}
