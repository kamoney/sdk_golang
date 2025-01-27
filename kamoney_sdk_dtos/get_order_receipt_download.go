package kamoney_sdk_dtos

type GetOrderReceiptDownloadRequestParams struct {
	Filename string `json:"file_name"`
}

type GetOrderReceiptDownloadRequestResponse struct {
	Common
	Data struct {
		Redirect string `json:"redirect"`
	} `json:"data,omitempty"`
}
