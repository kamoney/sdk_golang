package kamoney_sdk_dtos

// type ServicesOrderRequestParams struct{}

type ServicesOrderRequestResponse struct {
	Common
	Data struct {
		PaymentSlips struct {
			Code         string   `json:"code"`
			Picture      string   `json:"picture"`
			Name         string   `json:"name"`
			Maintenance  bool     `json:"maintenance"`
			LimitMin     int64    `json:"limit_min"`
			LimitMax     int64    `json:"limit_max"`
			FeeThreshold int64    `json:"fee_threshold"`
			FeeAmount    int64    `json:"fee_amount"`
			Automatic    bool     `json:"automatic"`
			AutomaticMax int64    `json:"automatic_max"`
			Info         []string `json:"info"`
		} `json:"payment_slips"`
		DirectTransfers struct {
			Code         string   `json:"code"`
			Picture      string   `json:"picture"`
			Name         string   `json:"name"`
			Maintenance  bool     `json:"maintenance"`
			LimitMin     int64    `json:"limit_min"`
			LimitMax     int64    `json:"limit_max"`
			FeeThreshold int64    `json:"fee_threshold"`
			FeeAmount    float64  `json:"fee_amount"`
			Automatic    bool     `json:"automatic"`
			AutomaticMax int64    `json:"automatic_max"`
			Info         []string `json:"info"`
		} `json:"direct_transfers"`
		Pix struct {
			Code         string   `json:"code"`
			Picture      string   `json:"picture"`
			Name         string   `json:"name"`
			Maintenance  bool     `json:"maintenance"`
			LimitMin     int64    `json:"limit_min"`
			LimitMax     int64    `json:"limit_max"`
			FeeThreshold int64    `json:"fee_threshold"`
			FeeAmount    int64    `json:"fee_amount"`
			Automatic    bool     `json:"automatic"`
			AutomaticMax int64    `json:"automatic_max"`
			Info         []string `json:"info"`
		} `json:"pix"`
		DigitalProducts struct {
			Code         string   `json:"code"`
			Picture      string   `json:"picture"`
			Name         string   `json:"name"`
			Maintenance  bool     `json:"maintenance"`
			LimitMin     int64    `json:"limit_min"`
			LimitMax     int64    `json:"limit_max"`
			FeeThreshold int64    `json:"fee_threshold"`
			FeeAmount    int64    `json:"fee_amount"`
			Automatic    bool     `json:"automatic"`
			AutomaticMax int64    `json:"automatic_max"`
			Info         []string `json:"info"`
		} `json:"digital_products"`
	} `json:"data"`
}
