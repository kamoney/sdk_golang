package kamoney_sdk_dtos

type GetAccountInfoRequestParams struct {
}

type GetAccountInfoRequestResponse struct {
	Common
	Data struct {
		Hash        string `json:"hash"`
		Email       string `json:"email"`
		Created     string `json:"created"`
		AccountType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"account_type"`
		Name                string `json:"name"`
		PersonalID          string `json:"personal_id"`
		DateOfBirth         string `json:"date_of_birth"`
		Whatsapp            string `json:"whatsapp"`
		Telegram            string `json:"telegram"`
		TFSEnabled          bool   `json:"tfs_enabled"`
		TFSActionPending    bool   `json:"tfs_action_pending"`
		AntiphishingEnabled bool   `json:"antiphishing_enabled"`
		Completed           bool   `json:"completed"`
		Active              bool   `json:"active"`
		Canceled            bool   `json:"canceled"`
		Merchant            bool   `json:"merchant"`
		AffiliateCode       string `json:"affiliate_code"`
	} `json:"data"`
}
