package kamoney_sdk_dtos

type ViewAntiPhishingRequestParams struct {
	Password string `json:"password"`
}

type ViewAntiPhishingRequestResponse struct {
	Common
	Data struct {
		Phrase string `json:"phrase"`
	} `json:"data"`
}
