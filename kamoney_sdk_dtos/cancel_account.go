package kamoney_sdk_dtos

type CancelAccountRequestParams struct {
	Password string `json:"password"`
	Terms    int    `json:"terms"`
	Nonce    string `json:"nonce"`
}

type CancelAccountRequestResponse struct {
	Common
}
