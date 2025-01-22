package kamoney_sdk_dtos

type GetAccountNotificationRequestParams struct {
	Page  int64  `json:"page"`
	Nonce string `json:"nonce"`
}

type GetAccountNotificationRequestResponse struct {
	Common
	Data struct {
		List []struct {
			ID      int64  `json:"id"`
			Created string `json:"created"`
			History string `json:"history"`
			Read    bool   `json:"read"`
		} `json:"list"`
		Unread int64 `json:"unread"`
	} `json:"data"`
}

type UpdateAccountNotificationReadAllRequestParams struct {
	Nonce string `json:"nonce"`
}

type UpdateAccountNotificationReadAllRequestResponse struct {
	Common
	Data struct {
		List []struct {
			ID      int64  `json:"id"`
			Created string `json:"created"`
			History string `json:"history"`
			Read    bool   `json:"read"`
			Title   string `json:"title"`
			Type    string `json:"type"`
		} `json:"list"`
		Unread int64 `json:"unread"`
	} `json:"data"`
}

type UpdateAccountNotificationReadIdRequestParams struct {
	ID    string `json:"id"`
	Nonce string `json:"nonce"`
}

type UpdateAccountNotificationReadIdRequestResponse struct {
	Common
	Data []struct {
		List []struct {
			ID      int64  `json:"id"`
			Created string `json:"created"`
			History string `json:"history"`
			Read    bool   `json:"read"`
			Title   string `json:"title"`
			Type    string `json:"type"`
		} `json:"list"`
		Unread int64 `json:"unread"`
	} `json:"data,omitempty"`
}
