package kamoney_sdk_dtos

type RegisterRequestParams struct {
	Email         string `json:"email"`
	AffiliateCode string `json:"affiliate_code"`
	Terms         bool   `json:"terms"`
}

type RegisterRequestResponse struct {
	Common
	Data []any `json:"data"`
}
