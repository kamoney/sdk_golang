package kamoney_sdk_dtos

type UtilsCityRequestResponse struct {
	Common
	Data []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}
