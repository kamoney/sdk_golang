package kamoney_sdk_dtos

type CancelAccountRequestParams struct {
	Password string `json:"password"`
	Terms    int64  `json:"terms"`
}

type CancelAccountRequestResponse struct {
	Common
}
