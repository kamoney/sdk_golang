package kamoney_sdk_dtos

type CreateMerchantRequestParams struct {
	Nonce          string  `json:"nonce"`
	Asset          string  `json:"asset"`
	Network        string  `json:"network"`
	Amount         float64 `json:"amount"`
	Email          string  `json:"email"`
	Callback       string  `json:"callback"`
	OrderID        string  `json:"order_id"`
	AdditionalInfo string  `json:"additional_info"`
	Redirect       string  `json:"redirect"`
}

type CreateMerchantRequestResponse struct {
	Common
	Data struct {
		ID             string  `json:"id"`
		Merchant       string  `json:"merchant"`
		OrderID        string  `json:"order_id"`
		AdditionalInfo string  `json:"additional_info"`
		Created        string  `json:"created"`
		Amount         float64 `json:"amount"`
		Quantity       string  `json:"quantity"`
		URI            string  `json:"uri"`
		Received       float64 `json:"received"`
		Asset          struct {
			Asset string `json:"asset"`
			Image string `json:"image"`
			Title string `json:"title"`
		} `json:"asset"`
		Address       string `json:"address"`
		BlockExplorer string `json:"block_explorer"`
		Memo          struct {
			Need bool   `json:"need"`
			Memo string `json:"memo"`
		} `json:"memo"`
		Status struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Img  string `json:"img"`
			Info string `json:"info"`
		} `json:"status"`
		Network struct {
			Name     string `json:"name"`
			Asset    string `json:"asset"`
			Protocol string `json:"protocol"`
		} `json:"network"`
	} `json:"data"`
}
