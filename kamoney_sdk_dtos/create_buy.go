package kamoney_sdk_dtos

type CreateBuyRequestParams struct {
	Nonce         string  `json:"nonce"`
	Asset         string  `json:"asset"`
	Network       string  `json:"network"`
	Amount        float64 `json:"amount"`
	WalletType    int64   `json:"wallet_type"`
	PaymentMethod string  `json:"payment_method"`
	Address       string  `json:"addr"`
	BaseCoin      string  `json:"base_coin,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
}

type CreateBuyRequestResponse struct {
	Common
	Data []struct {
		ID        string  `json:"id"`
		Created   string  `json:"created"`
		Amount    float64 `json:"amount"`
		Quotation struct {
			Value  float64 `json:"value"`
			Change bool    `json:"change"`
			Old    float64 `json:"old"`
			Time   struct {
				Start     string `json:"start"`
				ExpiredIn string `json:"expired_in"`
			} `json:"time"`
		} `json:"quotation"`
		QuantityTotal string `json:"quantity_total"`
		Quantity      string `json:"quantity"`
		Asset         struct {
			Asset string `json:"asset"`
			Image string `json:"image"`
			Title string `json:"title"`
		} `json:"asset"`
		Network struct {
			Name     string `json:"name"`
			Asset    string `json:"asset"`
			Protocol string `json:"protocol"`
		} `json:"network"`
		WalletType struct {
			Name string `json:"name"`
			Code int64  `json:"code"`
		} `json:"wallet_type"`
		Address       string `json:"address"`
		BlockExplorer string `json:"block_explorer"`
		Memo          struct {
			Need bool   `json:"need"`
			Memo string `json:"memo"`
		} `json:"memo"`
		Fee    float64 `json:"fee"`
		Status struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Img  string `json:"img"`
		} `json:"status"`
		PaymentMethod struct {
			DataPayment struct {
				Created      string  `json:"created"`
				Amount       float64 `json:"amount"`
				PayTime      string  `json:"pay_time"`
				PayTimeLimit string  `json:"pay_time_limit"`
				QRCode       string  `json:"qrcode"`
				Status       struct {
					Name string `json:"name"`
					Code string `json:"code"`
					Img  string `json:"img"`
				} `json:"status"`
			} `json:"data_payment"`
		} `json:"payment_method"`
	} `json:"data"`
}
