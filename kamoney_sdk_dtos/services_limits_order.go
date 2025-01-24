package kamoney_sdk_dtos

type GetServicesLimitsOrderRequestParams struct {
}

type GetServicesLimitsOrderRequestResponse struct {
	Data struct {
		PaymentSlips    Service `json:"payment_slips"`
		DirectTransfers Service `json:"direct_transfers"`
		Pix             Service `json:"pix"`
		DigitalProducts Service `json:"digital_products"`
	} `json:"data"`
}

type Service struct {
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
	Limits       struct {
		Owner int64 `json:"owner"`
		Other int64 `json:"other"`
	} `json:"limits"`
}
