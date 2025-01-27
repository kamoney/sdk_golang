package kamoney_sdk_dtos

type GetBuyPrivateKeyRequestParams struct {
	Password string `json:"password"`
	TFA      string `json:"tfa"`
}

type GetBuyPrivateKeyRequestResponse struct {
	Common
	Data struct {
		Private string `json:"private"`
	} `json:"data"`
}
