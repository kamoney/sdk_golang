package kamoney_sdk_dtos

type CreateAPIRequestParams struct {
	Nonce string `json:"nonce"`
}

type CreateAPIRequestResponse struct {
	Common
	Data struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	} `json:"data"`
}
