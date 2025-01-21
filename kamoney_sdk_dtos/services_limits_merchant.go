package kamoney_sdk_dtos

type GetServicesLimitsMerchantRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetServicesLimitsMerchantRequestResponse struct {
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
			Owner int `json:"owner"`
			Other int `json:"other"`
		} `json:"limits"`
	} `json:"data"`
}
