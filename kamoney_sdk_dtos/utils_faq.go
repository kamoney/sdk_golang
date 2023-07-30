package kamoney_sdk_dtos

type UtilsFaqRequestResponse struct {
	Common
	Data []struct {
		ID      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"data"`
}
