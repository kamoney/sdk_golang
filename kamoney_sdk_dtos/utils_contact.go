package kamoney_sdk_dtos

type UtilsContactRequestParams struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type UtilsContactRequestResponse struct {
	Common
}
