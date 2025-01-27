package kamoney_sdk_dtos

type ListAPIRequestParams struct {
	Page string `json:"page,omitempty"`
}

type ListAPIRequestResponse struct {
	Common
	Data []struct {
		ID              int64  `json:"id"`
		Public          string `json:"public"`
		Created         string `json:"created"`
		IP              string `json:"ip"`
		EnabledServices int64  `json:"enabled_services"`
		EnabledMerchant int64  `json:"enabled_merchant"`
	} `json:"data"`
}
