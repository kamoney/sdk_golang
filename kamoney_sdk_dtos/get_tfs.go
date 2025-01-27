package kamoney_sdk_dtos

type GetTFSRequestParams struct {
}

type GetTFSRequestResponse struct {
	Common
	Data struct {
		Enabled bool   `json:"enabled"`
		Secret  string `json:"secret"`
	} `json:"data"`
}
