package kamoney_sdk_dtos

type ListOrderRequestParams struct {
	Nonce  string `json:"nonce"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Search string `json:"search"`
	Status string `json:"status"`
}

type ListOrderRequestResponse struct {
	Common
	Data []struct {
		ID       string  `json:"id"`
		Created  string  `json:"created"`
		Amount   float64 `json:"amount"`
		Quantity string  `json:"quantity"`
		Received float64 `json:"received"`
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
	} `json:"data"`
}
