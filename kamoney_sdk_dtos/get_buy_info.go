package kamoney_sdk_dtos

type GetBuyInfoRequestParams struct{}

type GetBuyInfoRequestResponse struct {
	Common
	Data struct {
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
		Quantity string `json:"quantity"`
		Asset    struct {
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
			Info string `json:"info"`
		} `json:"status"`
		PaymentMethod struct {
			Created      string  `json:"created"`
			Amount       float64 `json:"amount"`
			PayTime      string  `json:"pay_time"`
			PayTimeLimit string  `json:"pay_time_limit"`
			QRCode       string  `json:"qrcode"`
			Status       struct {
				Name string `json:"name"`
				Code string `json:"code"`
				Img  string `json:"img"`
				Info string `json:"info"`
			} `json:"status"`
		} `json:"payment_method"`
	} `json:"data"`
}
