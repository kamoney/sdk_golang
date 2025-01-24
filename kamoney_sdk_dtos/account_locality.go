package kamoney_sdk_dtos

type AccountLocalityRequestParams struct {
	Zipcode      string `json:"zipcode"`
	Street       string `json:"street"`
	Number       int64  `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         int64  `json:"city"`
	State        int64  `json:"state"`
	Nonce        string `json:"nonce"`
}

type AccountLocalityRequestResponse struct {
	Common
	Data struct {
		Street       string `json:"street"`
		Number       string `json:"number"`
		Complement   string `json:"complement"`
		Neighborhood string `json:"neighborhood"`
		Zipcode      string `json:"zipcode"`
		City         struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
			UF   string `json:"uf"`
		} `json:"state"`
		Country struct {
			ID       int64  `json:"id"`
			Name     string `json:"name"`
			Initials string `json:"initials"`
		} `json:"country"`
	} `json:"data"`
}

type GetAccountLocalityRequestParams struct {
}

type GetAccountLocalityRequestResponse struct {
	Common
	Data struct {
		Street       string `json:"street"`
		Number       string `json:"number"`
		Complement   string `json:"complement"`
		Neighborhood string `json:"neighborhood"`
		Zipcode      string `json:"zipcode"`
		City         struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
			UF   string `json:"uf"`
		} `json:"state"`
		Country struct {
			ID       int64  `json:"id"`
			Name     string `json:"name"`
			Initials string `json:"initials"`
		} `json:"country"`
	} `json:"data"`
}
