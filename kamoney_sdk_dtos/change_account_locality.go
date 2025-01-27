package kamoney_sdk_dtos

type ChangeAccountLocalityRequestParams struct {
	Zipcode      string `json:"zipcode"`
	Street       string `json:"street"`
	Number       int64  `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         int64  `json:"city"`
	Nonce        string `json:"nonce"`
}

type ChangeAccountLocalityRequestResponse struct {
	Common
}
