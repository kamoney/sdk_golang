package kamoney_sdk_dtos

type DeleteAPIRequestParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Tfa      int64  `json:"tfa"`
}

type DeleteAPIRequestResponse struct {
	Common
}
