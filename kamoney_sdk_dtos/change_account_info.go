package kamoney_sdk_dtos

type ChangeAccountInfoRequestParams struct {
	Name        string `json:"name"`
	PersonalID  string `json:"personal_id"`
	DateOfBirth string `json:"date_of_birth"`
}

type ChangeAccountInfoRequestResponse struct {
	Common
}
