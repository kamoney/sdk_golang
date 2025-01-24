package kamoney_sdk_dtos

type ViewTfsRequestParams struct {
}

type ViewTfsRequestResponse struct {
	Common
	Data struct {
		Enabled bool   `json:"enabled"`
		Secret  string `json:"secret"`
	} `json:"data"`
}
