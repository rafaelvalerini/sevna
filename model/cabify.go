package model

type ResponseCabify []struct {
	VehicleType struct {
		ID string `json:"_id"`
		Name string `json:"name"`
		ShortName string `json:"short_name"`
		Description string `json:"description"`
		Icons struct {
			Regular string `json:"regular"`
		} `json:"icons"`
		ReservedOnly bool `json:"reserved_only"`
		AsapOnly bool `json:"asap_only"`
		Currency string `json:"currency"`
		Icon string `json:"icon"`
	} `json:"vehicle_type"`
	TotalPrice float64 `json:"total_price"`
	FormattedPrice string `json:"price_formatted"`
	Currency string `json:"currency"`
	CurrencySymbol string `json:"currency_symbol"`
}

type RequestCabify struct {
	Stop []Stop `json:"stops"`
	StartAt string `json:"start_at"`
}

type Stop struct{
	Loc []float64 `json:"loc"`
}