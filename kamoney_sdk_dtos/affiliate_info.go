package kamoney_sdk_dtos

type GetAffiliateInfoRequestParams struct {
	Nonce string `json:"nonce"`
}

type GetAffiliateInfoRequestResponse struct {
	Common
	Data struct {
		AffiliateCode string `json:"affiliate_code"`
		Report        struct {
			Indicated int `json:"indicated"`
			Orders    int `json:"orders"`
			Total     int `json:"total"`
		} `json:"report"`
		List []struct {
			Name    string `json:"name"`
			Email   string `json:"email"`
			Created string `json:"created"`
		} `json:"list"`
	} `json:"data,omitempty"`
}
