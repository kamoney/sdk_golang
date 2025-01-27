package kamoney_sdk_dtos

type ChangeAccountContactRequestParams struct {
	Whatsapp string `json:"whatsapp"`
	Telegram string `json:"telegram"`
}

type ChangeAccountContactRequestResponse struct {
	Common
}
