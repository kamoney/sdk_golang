package kamoney_sdk_dtos

type GetAPISecretRequestParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Tfa      int64  `json:"tfa"` // in case tfa is active
	Nonce    string `json:"nonce"`
}

type GetAPISecretRequestResponse struct {
	Common
	Data struct {
		Secret string `json:"secret"`
	} `json:"data"`
}
