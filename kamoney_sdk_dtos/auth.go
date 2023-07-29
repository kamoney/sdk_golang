package kamoney_sdk_dtos

type AuthRequestParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRequestResponse struct {
	Common
	Data struct {
		TfsEnabled bool   `json:"tfs_enabled"`
		TokenExp   string `json:"token_exp"`
		Token      string `json:"token"`
	} `json:"data"`
}
