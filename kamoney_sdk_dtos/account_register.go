package kamoney_sdk_dtos

type AccountRegisterRequestParams struct {
	Email         string `json:"email"`
	AffiliateCode string `json:"affiliate_code"`
	Terms         bool   `json:"terms"`
}

type AccountRegisterRequestResponse struct {
	Common
	Data []any `json:"data"`
}
