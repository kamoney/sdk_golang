package kamoney_sdk_dtos

type GetAccountKycRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetAccountKycRequestResponse struct {
	Common
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}
