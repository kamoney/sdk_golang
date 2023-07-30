package kamoney_sdk_dtos

type UtilsCurrencyNetworkRequestResponse struct {
	Common
	Data []struct {
		Name         string  `json:"name"`
		Symbol       string  `json:"symbol"`
		Protocol     string  `json:"protocol"`
		Memo         bool    `json:"memo"`
		Enabled      bool    `json:"enabled"`
		EnabledSell  bool    `json:"enabled_sell"`
		EnabledBuy   bool    `json:"enabled_buy"`
		ExplorerAddr string  `json:"explorer_addr"`
		ExplorerTx   string  `json:"explorer_tx"`
		FeeWithdraw  float64 `json:"fee_withdraw"`
	} `json:"data"`
}
