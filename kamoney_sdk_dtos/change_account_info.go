package kamoney_sdk_dtos

type ChangeAccountInfoRequestParams struct {
	Name        string `json:"name"`
	PersonalID  string `json:"personal_id"`
	DateOfBirth string `json:"date_of_birth"`
	Nonce       string `json:"nonce"`
}

type ChangeAccountInfoRequestResponse struct {
	Common
}
