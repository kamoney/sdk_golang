package kamoney_sdk_dtos

type UtilsNotificationRequestResponse struct {
	Common
	Data []struct {
		ID          int64  `json:"id"`
		Category    string `json:"category"`
		Subcategory string `json:"subcategory"`
		Message     string `json:"message"`
	} `json:"data"`
}
