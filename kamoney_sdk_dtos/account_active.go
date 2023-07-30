package kamoney_sdk_dtos

type AccountActiveRequestParams struct {
	Email    string `json:"email"`
	Code     int64  `json:"code"`
	Password string `json:"password"`
}

type AccountActiveRequestResponse struct {
	Common
}
