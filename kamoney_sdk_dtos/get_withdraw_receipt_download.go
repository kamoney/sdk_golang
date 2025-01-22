package kamoney_sdk_dtos

type GetWithdrawReceiptDownloadRequestParams struct {
	Nonce    string `json:"nonce"`
	ID       string `json:"id"`
	FileName string `json:"filename"`
}

type GetWithdrawReceiptDownloadRequestResponse struct {
	Common
	Data []struct {
		Redirect string `json:"redirect"`
	} `json:"data"`
}
