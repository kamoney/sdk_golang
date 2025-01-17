package kamoney_sdk_dtos

type AccountInfoRequestParams struct {
	Name        string `json:"name"`
	PersonalID  string `json:"personal_id"`
	DateOfBirth string `json:"date_of_birth"`
}

type AccountInfoRequestResponse struct {
	Common
}
