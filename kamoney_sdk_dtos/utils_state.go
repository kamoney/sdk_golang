package kamoney_sdk_dtos

type UtilsStateRequestResponse struct {
	Common
	Data []struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Initials string `json:"initials"`
	} `json:"data"`
}
