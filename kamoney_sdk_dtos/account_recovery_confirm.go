package kamoney_sdk_dtos

type AccountRecoveryConfirmRequestParams struct {
	Email    string `json:"email"`
	Code     int64  `json:"code"`
	Password string `json:"password"`
}

type AccountRecoveryConfirmRequestResponse struct {
	Common
}
