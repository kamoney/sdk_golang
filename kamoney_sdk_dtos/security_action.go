package kamoney_sdk_dtos

type SecurityActionRequestParams struct {
	Code  string `json:"code"`
	Nonce string `json:"nonce"`
}

type SecurityActionRequestResponse struct {
	Common
}
