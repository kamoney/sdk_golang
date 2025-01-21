package kamoney_sdk_dtos

type AccountContactRequestParams struct {
	Whatsapp string `json:"whatsapp"`
	Telegram string `json:"telegram"`
	Nonce    string `json:"nonce"`
}

type AccountContactRequestResponse struct {
	Common
}
