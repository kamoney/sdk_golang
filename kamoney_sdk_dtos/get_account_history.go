package kamoney_sdk_dtos

type GetAccountHistoryRequestParams struct {
	Page   int64  `json:"page"`
	Search string `json:"search"`
	Date   string `json:"date"`
}

type GetAccountHistoryRequestResponse struct {
	Common
	Data []struct {
		ID      int64  `json:"id"`
		Created string `json:"created"`
		History string `json:"history"`
	} `json:"data"`
}
