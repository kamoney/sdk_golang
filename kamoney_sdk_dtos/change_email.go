package kamoney_sdk_dtos

type ChangeEmailRequestParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Tfa      string `json:"tfa"`
	Nonce    string `json:"nonce"`
}

type ChangeEmailRequestResponse struct {
	Common
}
