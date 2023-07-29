package kamoney_sdk_dtos

type RecoveryConfirmRequestParams struct {
	Email    string `json:"email"`
	Code     int64  `json:"code"`
	Password string `json:"password"`
}

type RecoveryConfirmRequestResponse struct {
	Common
}
