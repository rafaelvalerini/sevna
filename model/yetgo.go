package model

type ResponseYetGo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		MinArrivalTime float64 `json:"min_arrival_time"`
		TripDistance   float64 `json:"trip_distance"`
		TripDuration   float64 `json:"trip_duration"`
		CurrencyCode   string  `json:"currency_code"`
		Currency       string  `json:"currency"`
		TripMinPrice   float64 `json:"trip_min_price"`
		TripMaxPrice   float64 `json:"trip_max_price"`
		Categories     []struct {
			DisplayName            string  `json:"display_name"`
			SubcategoryDisplayName string  `json:"subcategory_display_name"`
			EstimatedPrice         string  `json:"estimated_price"`
			MinDistance            float64 `json:"min_distance"`
		} `json:"categories"`
	} `json:"data"`
}
