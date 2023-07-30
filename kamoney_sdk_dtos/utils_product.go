package kamoney_sdk_dtos

type UtilsProductRequestResponse struct {
	Common
	Data []struct {
		ID        int64   `json:"id"`
		Name      string  `json:"name"`
		Label     string  `json:"label"`
		Price     float64 `json:"price"`
		Min       float64 `json:"min"`
		Logo      string  `json:"logo"`
		Pricebase float64 `json:"pricebase"`
		Max       float64 `json:"max"`
		Maxbuy    float64 `json:"maxbuy"`
	} `json:"data"`
}
