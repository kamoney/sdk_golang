package kamoney_sdk_dtos

type CreateWithdrawRequestParams struct {
	Nonce  string `json:"nonce"`
	Asset  string `json:"asset"`
	Type   string `json:"type"`
	Key    string `json:"key"`
	Amount int64  `json:"amount"`
	Tfa    string `json:"tfa"`
}

type CreateWithdrawRequestResponse struct {
	Common
	Data struct {
		ID      string `json:"id"`
		Created string `json:"created"`
		Symbol  string `json:"symbol"`
		Amount  int64  `json:"amount"`
		Fee     int64  `json:"fee"`
		Total   int64  `json:"total"`
		Type    string `json:"type"`
		Owner   struct {
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
		Pix struct {
			Key  string `json:"key"`
			Type string `json:"type"`
		} `json:"pix"`
		Modality int64 `json:"modality"`
	} `json:"data"`
}
