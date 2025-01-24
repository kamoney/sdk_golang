package kamoney_sdk_dtos

type GetWithdrawReceiptRequestParams struct {
	ID string `json:"id"`
}

type GetWithdrawReceiptRequestResponse struct {
	Common
	Data []struct {
		ID       int64  `json:"id"`
		Created  string `json:"created"`
		FileName string `json:"file_name"`
	} `json:"data"`
}
