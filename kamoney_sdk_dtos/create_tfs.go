package kamoney_sdk_dtos

type CreateTfsRequestParams struct {
	Password string `json:"password"`
	Nonce    string `json:"nonce"`
}

type CreateTfsRequestResponse struct {
	Common
}
