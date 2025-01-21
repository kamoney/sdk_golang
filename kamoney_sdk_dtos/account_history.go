package kamoney_sdk_dtos

type GetAccountHistoryRequestParams struct {
	Page  int    `json:"page"`
	Date  int    `json:"date"`
	Nonce string `json:"nonce"`
}

type GetAccountHistoryRequestResponse struct {
	Common
	Data []struct {
		ID      int    `json:"id"`
		Created string `json:"created"`
		History string `json:"history"`
	} `json:"data"`
}
