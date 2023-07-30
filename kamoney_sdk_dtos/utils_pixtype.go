package kamoney_sdk_dtos

type UtilsPixTypeRequestResponse struct {
	Common
	Data []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"data"`
}
