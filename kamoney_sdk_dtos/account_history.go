package kamoney_sdk_dtos

type GetAccountHistoryRequestParams struct {
	Page  int64  `json:"page"`
	Date  int64  `json:"date"`
	Nonce string `json:"nonce"`
}

type GetAccountHistoryRequestResponse struct {
	Common
	Data []struct {
		ID      int64  `json:"id"`
		Created string `json:"created"`
		History string `json:"history"`
	} `json:"data"`
}
