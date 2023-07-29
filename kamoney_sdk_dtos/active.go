package kamoney_sdk_dtos

type ActiveRequestParams struct {
	Email    string `json:"email"`
	Code     int64  `json:"code"`
	Password string `json:"password"`
}

type ActiveRequestResponse struct {
	Common
}
