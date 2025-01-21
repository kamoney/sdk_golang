package kamoney_sdk_dtos

type GetRewardRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetRewardRequestResponse struct {
	Common
	Data struct {
		Affiliate int `json:"affiliate"`
		Points    int `json:"points"`
	} `json:"data"`
}
