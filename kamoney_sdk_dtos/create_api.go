package kamoney_sdk_dtos

type CreateAPIRequestParams struct {
}

type CreateAPIRequestResponse struct {
	Common
	Data struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	} `json:"data"`
}
