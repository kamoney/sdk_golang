package kamoney_sdk_dtos

type UtilsBankRequestResponse struct {
	Common
	Data []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"data"`
}
