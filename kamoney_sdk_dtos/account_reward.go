package kamoney_sdk_dtos

type GetRewardRequestParams struct {
}

type GetRewardRequestResponse struct {
	Common
	Data struct {
		Affiliate int64 `json:"affiliate"`
		Points    int64 `json:"points"`
	} `json:"data"`
}
