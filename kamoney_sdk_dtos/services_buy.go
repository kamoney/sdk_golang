package kamoney_sdk_dtos

type ServicesBuyRequestResponse struct {
	Common
	Data struct {
		Code         string   `json:"code"`
		Name         string   `json:"name"`
		Maintenance  bool     `json:"maintenance"`
		LimitMin     float64  `json:"limit_min"`
		LimitMax     float64  `json:"limit_max"`
		FeeThreshold float64  `json:"fee_threshold"`
		FeeAmount    float64  `json:"fee_amount"`
		Automatic    bool     `json:"automatic"`
		AutomaticMax float64  `json:"automatic_max"`
		Info         []string `json:"info"`
	} `json:"data"`
}
