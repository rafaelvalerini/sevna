package model

type ResponseTime struct {
	Times    []Times  `json:"times"`
}

type Times struct {
	LocalizedDisplayName    string  `json:"localized_display_name"`
	Estimate       			int  `json:"estimate"`
	DisplayName       		string  `json:"display_name"`
	ProductId       		string  `json:"product_id"`
}

type ResponsePrices struct {
	Prices []struct {
		ProductID string `json:"product_id"`
		CurrencyCode string `json:"currency_code"`
		DisplayName string `json:"display_name"`
		Estimate string `json:"estimate"`
		LowEstimate int `json:"low_estimate"`
		HighEstimate int `json:"high_estimate"`
		SurgeMultiplier float64 `json:"surge_multiplier"`
		Duration int `json:"duration"`
		Distance float64 `json:"distance"`
	} `json:"prices"`
}

type ResponseProduct struct {
	Products []struct {
		Capacity int `json:"capacity"`
		Description string `json:"description"`
		PriceDetails struct {
			DistanceUnit string `json:"distance_unit"`
			CostPerMinute float64 `json:"cost_per_minute"`
			ServiceFees []struct {
				Fee float64 `json:"fee"`
				Name string `json:"name"`
			} `json:"service_fees"`
			Minimum float64 `json:"minimum"`
			CostPerDistance float64 `json:"cost_per_distance"`
			Base float64 `json:"base"`
			CancellationFee float64 `json:"cancellation_fee"`
			CurrencyCode string `json:"currency_code"`
		} `json:"price_details"`
		Image string `json:"image"`
		DisplayName string `json:"display_name"`
		ProductID string `json:"product_id"`
		Shared bool `json:"shared"`
	} `json:"products"`
}