package kamoney_sdk_dtos

type GetOrderInfoRequestParams struct {
	ID string `json:"id"`
}

type GetOrderInfoRequestResponse struct {
	Common
	Data struct {
		ID       string  `json:"id"`
		Created  string  `json:"created"`
		Amount   float64 `json:"amount"`
		Quantity string  `json:"quantity"`
		URI      string  `json:"uri"`
		Received int64   `json:"received"`
		Asset    struct {
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
		Services struct {
			PaymentSlips    []interface{} `json:"payment_slips"`
			DirectTransfers []interface{} `json:"direct_transfers"`
			DigitalProducts []interface{} `json:"digital_products"`
			Pix             []struct {
				Key    string `json:"key"`
				Type   string `json:"type"`
				Amount int    `json:"amount"`
				Fee    int    `json:"fee"`
				Total  int    `json:"total"`
				Status struct {
					Name string `json:"name"`
					Code string `json:"code"`
					Img  string `json:"img"`
				} `json:"status"`
			} `json:"pix"`
		} `json:"services"`
	} `json:"data"`
}
