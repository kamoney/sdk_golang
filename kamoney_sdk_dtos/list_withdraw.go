package kamoney_sdk_dtos

type ListWithdrawRequestParams struct {
}

type ListWithdrawRequestResponse struct {
	Common
	Data []struct {
		ID       string  `json:"id"`
		Created  string  `json:"created"`
		Symbol   string  `json:"symbol"`
		Amount   float64 `json:"amount"`
		Fee      float64 `json:"fee"`
		Total    float64 `json:"total"`
		Modality int64   `json:"modality"`
		BankName string  `json:"bank_name"`
		Owner    struct {
			Name       string `json:"name"`
			PersonalID string `json:"personal_id"`
		} `json:"owner"`
		Account struct {
			Number string `json:"number"`
			Type   string `json:"type"`
			Agency string `json:"agency"`
		} `json:"account"`
		Status struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Img  string `json:"img"`
			Info string `json:"info"`
		} `json:"status"`
		Receipt struct {
			Attached bool   `json:"attached"`
			FileName string `json:"file_name"`
		} `json:"receipt"`
	} `json:"data"`
}
