package kamoney_sdk_dtos

type GetAccountKycRequestParams struct {
}

type GetAccountKycRequestResponse struct {
	Common
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}
