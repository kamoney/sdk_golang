package kamoney_sdk_dtos

type AccountLocalityRequestParams struct{}

type AccountLocalityRequestResponse struct {
	Common
	Data struct {
		Street       string `json:"street"`
		Number       string `json:"number"`
		Complement   string `json:"complement"`
		Neighborhood string `json:"neighborhood"`
		Zipcode      string `json:"zipcode"`
		City         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			UF   string `json:"uf"`
		} `json:"state"`
		Country struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Initials string `json:"initials"`
		} `json:"country"`
	} `json:"data"`
}