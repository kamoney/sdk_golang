package kamoney_sdk_dtos

type ListOrderReceiptRequestParams struct {
	Nonce  string `json:"nonce"`
	Begin  string `json:"begin"`
	End    string `json:"end"`
	Search string `json:"search"`
}

type ListOrderReceiptRequestResponse struct {
	Common
	Data []struct {
		ID       int    `json:"id"`
		Created  string `json:"created"`
		Order    int    `json:"order"`
		Hash     string `json:"hash"`
		FileName string `json:"file_name"`
	} `json:"data"`
}
