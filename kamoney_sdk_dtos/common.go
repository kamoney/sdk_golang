package kamoney_sdk_dtos

type Common struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Error   string `json:"error"`
	Code    string `json:"code"`
}
