package kamoney_sdk_dtos

type ChangePasswordRequestParams struct {
	Password        string `json:"password"`
	PasswordNew     string `json:"password_new"`
	PasswordConfirm string `json:"password_confirm"`
	Nonce           string `json:"nonce"`
}

type ChangePasswordRequestResponse struct {
	Common
}
