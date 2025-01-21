package kamoney_sdk_dtos

type ListAPIsRequestParams struct {
	Page  string `json:"page,omitempty"`
	Nonce string `json:"nonce"`
}

type ListAPIsRequestResponse struct {
	Common
	Data []struct {
		ID              int    `json:"id"`
		Public          string `json:"public"`
		Created         string `json:"created"`
		IP              string `json:"ip"`
		EnabledServices int    `json:"enabled_services"`
		EnabledMerchant int    `json:"enabled_merchant"`
	} `json:"data"`
}
