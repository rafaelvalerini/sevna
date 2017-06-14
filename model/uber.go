package model

type ResponseTime struct {
	Times []Times `json:"times"`
}

type Times struct {
	LocalizedDisplayName string `json:"localized_display_name"`
	Estimate             int    `json:"estimate"`
	DisplayName          string `json:"display_name"`
	SortDescription      string `json:"short_description"`
	ProductId            string `json:"product_id"`
}

type ResponsePrices struct {
	Prices []struct {
		ProductID       string  `json:"product_id"`
		CurrencyCode    string  `json:"currency_code"`
		DisplayName     string  `json:"display_name"`
		Estimate        string  `json:"estimate"`
		LowEstimate     int     `json:"low_estimate"`
		HighEstimate    int     `json:"high_estimate"`
		SurgeMultiplier float64 `json:"surge_multiplier"`
		Duration        int     `json:"duration"`
		Distance        float64 `json:"distance"`
	} `json:"prices"`
}

type ResponseProduct struct {
	Products []struct {
		Capacity     int    `json:"capacity"`
		Description  string `json:"description"`
		PriceDetails struct {
			DistanceUnit  string  `json:"distance_unit"`
			CostPerMinute float64 `json:"cost_per_minute"`
			ServiceFees   []struct {
				Fee  float64 `json:"fee"`
				Name string  `json:"name"`
			} `json:"service_fees"`
			Minimum         float64 `json:"minimum"`
			CostPerDistance float64 `json:"cost_per_distance"`
			Base            float64 `json:"base"`
			CancellationFee float64 `json:"cancellation_fee"`
			CurrencyCode    string  `json:"currency_code"`
		} `json:"price_details"`
		Image           string `json:"image"`
		DisplayName     string `json:"display_name"`
		SortDescription string `json:"short_description"`
		ProductID       string `json:"product_id"`
		Shared          bool   `json:"shared"`
	} `json:"products"`
}

type ResponseEstimateV12 struct {
	Fare struct {
		Value        float64 `json:"value"`
		FareID       string  `json:"fare_id"`
		ExpiresAt    int     `json:"expires_at"`
		Display      string  `json:"display"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"fare"`
	Trip struct {
		DistanceUnit     string  `json:"distance_unit"`
		DurationEstimate int     `json:"duration_estimate"`
		DistanceEstimate float64 `json:"distance_estimate"`
	} `json:"trip"`
	PickupEstimate int `json:"pickup_estimate"`
}

type ResponseProductV12 struct {
	Products []ProductUber `json:"products"`
}

type ProductUber struct {
	UpfrontFareEnabled bool   `json:"upfront_fare_enabled"`
	Capacity           int    `json:"capacity"`
	ProductID          string `json:"product_id"`
	Image              string `json:"image"`
	CashEnabled        bool   `json:"cash_enabled"`
	Shared             bool   `json:"shared"`
	ShortDescription   string `json:"short_description"`
	DisplayName        string `json:"display_name"`
	ProductGroup       string `json:"product_group"`
	Description        string `json:"description"`
}

type RequestEstimateV12 struct {
	ProductID      string  `json:"product_id"`
	StartLatitude  float64 `json:"start_latitude"`
	StartLongitude float64 `json:"start_longitude"`
	EndLatitude    float64 `json:"end_latitude"`
	EndLongitude   float64 `json:"end_longitude"`
	SeatCount      string  `json:"seat_count"`
}

type TokenUber struct {
	Id          int
	Active      int
	TokenBearer string
}
