package kamoney_sdk_dtos

type AccountRecoveryRequestParams struct {
	Email string `json:"email"`
}

type AccountRecoveryRequestResponse struct {
	Common
}
