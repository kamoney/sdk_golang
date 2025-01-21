package kamoney_sdk_dtos

type ViewAntiPhishingRequestParams struct {
	Password string `json:"password"`
	Nonce    string `json:"nonce"`
}

type ViewAntiPhishingRequestResponse struct {
	Common
	Data struct {
		Phrase string `json:"phrase"`
	} `json:"data"`
}
