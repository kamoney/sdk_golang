package kamoney_sdk_dtos

type AccountContactRequestParams struct {
	Whatsapp string `json:"whatsapp"`
	Telegram string `json:"telegram"`
}

type AccountContactRequestResponse struct {
	Common
}
