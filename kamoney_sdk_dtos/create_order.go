package kamoney_sdk_dtos

type CreateOrderRequestParams struct {
	Nonce        string `json:"nonce,omitempty"`
	Asset        string `json:"asset"`
	Network      string `json:"network"`
	Coupon       string `json:"coupon,omitempty"`
	PaymentSlips []struct {
		Barcode     string  `json:"barcode"`
		Institution string  `json:"institution"`
		Amount      float64 `json:"amount"`
		DueDate     string  `json:"due_date"`
		// PersonalID  string  `json:"personal_id"`
		// Owner       string  `json:"owner"`
	} `json:"payment_slips,omitempty"`
	DirectTransfers []struct {
		AccountType   string  `json:"account_type"`
		BankID        int64   `json:"bank_id"`
		Agency        string  `json:"agency"`
		AccountNumber string  `json:"account_number"`
		PersonalID    string  `json:"personal_id"`
		Owner         string  `json:"owner"`
		Amount        float64 `json:"amount"`
	} `json:"direct_transfers,omitempty"`
	DigitalProducts []struct {
		ProductID int64 `json:"product_id"`
		Quantity  int64 `json:"quantity"`
	} `json:"digital_products,omitempty"`
	Pix []struct {
		Type   string  `json:"type"`
		Key    string  `json:"key"`
		Amount float64 `json:"amount"`
	} `json:"pix,omitempty"`
}

type CreateOrderRequestResponse struct {
	Common
	Data struct {
		ID       string  `json:"id"`
		Created  string  `json:"created"`
		Amount   float64 `json:"amount"`
		Quantity string  `json:"quantity"`
		URI      string  `json:"uri"`
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
		Network struct {
			Name     string `json:"name"`
			Asset    string `json:"asset"`
			Protocol string `json:"protocol"`
		} `json:"network"`
		Services struct {
			PaymentSlips []struct {
				Barcode     string  `json:"barcode"`
				Institution string  `json:"institution"`
				DueDate     string  `json:"due_date"`
				PersonalID  string  `json:"personal_id"`
				Owner       string  `json:"owner"`
				Amount      float64 `json:"amount"`
				Fee         float64 `json:"fee"`
				Total       float64 `json:"total"`
				Status      struct {
					Name string `json:"name"`
					Code string `json:"code"`
					Img  string `json:"img"`
				} `json:"status"`
			} `json:"payment_slips"`
			MobileRecharges []interface{} `json:"mobile_recharges"`
			DirectTransfers []interface{} `json:"direct_transfers"`
			DigitalProducts []interface{} `json:"digital_products"`
			Pix             []interface{} `json:"pix"`
		} `json:"services"`
	} `json:"data"`
}
