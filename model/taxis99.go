package model

type Response99Taxi struct {
	Estimates []struct {
		CategoryID string `json:"categoryId"`
		CurrencyCode string `json:"currencyCode"`
		LowerFare string `json:"lowerFare"`
		UpperFare string `json:"upperFare"`
	} `json:"estimates"`
	PricingEstimateID int `json:"pricingEstimateId"`
}

type Request99Taxi struct {
	DistanceInMeters int64 `json:"distanceInMeters"`
	DropoffLatitude float64 `json:"dropoffLatitude"`
	DropoffLongitude float64 `json:"dropoffLongitude"`
	PickupLatitude float64 `json:"pickupLatitude"`
	PickupLongitude float64 `json:"pickupLongitude"`
	TimeInSeconds int64 `json:"timeInSeconds"`
}