package kamoney_sdk_dtos

type UtilsCurrencyRequestResponse struct {
	Common
	Data []struct {
		Image           string  `json:"image"`
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol"`
		Asset           string  `json:"asset"`
		Quotation       float64 `json:"quotation"`
		Decimals        float64 `json:"decimals"`
		Fluctuation     float64 `json:"fluctuation"`
		Maintenance     bool    `json:"maintenance"`
		EnabledOrder    bool    `json:"enabled_order"`
		EnabledMerchant bool    `json:"enabled_merchant"`
		EnabledBuy      bool    `json:"enabled_buy"`
		EnabledDeposit  bool    `json:"enabled_deposit"`
		EnabledWithdraw bool    `json:"enabled_withdraw"`
		EnabledMobile   bool    `json:"enabled_mobile"`
	} `json:"data"`
}
