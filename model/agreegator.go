package model


type Aggregator struct{
	Id string `json:"id"`
	Start Position  `json:"start"`
	End Position  `json:"end"`
	Players []Player `json:"records"` 
}

type RequestAggregator struct{
	Start Position  `json:"start"`
	End Position  `json:"end"`
	Device Device `json:"device"`
	Distance int64 `json:"distance"`
	Duration int64 `json:"duration"`
}

type Device struct{
	IdDevice string `json:"id"`
	OperationSystem string `json:"operation_system"`
	OperationSystemVersion string `json:"operation_system_version"`
	Device string `json:"device"`
	TypeConnection string `json:"type_connection"`
}


type Position struct{
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Address string `json:"address"`
	District string `json:"district"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zipcode"`
}

type Player struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Modality Modality `json:"modality,omitempty"`
	WaitingTime int `json:"waiting_time,omitempty"`
	Price string `json:"price,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Multiplier float64 `json:"multiplier,omitempty"`
	Modalities []Modality `json:"modalities,omitempty"`
}

type Modality struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Promotion Promotion `json:"promotion,omitempty"`
	PriceKm float64 `json:"price_km,omitempty"`
	TimeKm int `json:"time_km,omitempty"`
	ModalityCoverage []Coverage `json:"coverages,omitempty"`
	PriceTime float64 `json:"price_time,omitempty"`
	PriceBase float64 `json:"price_base,omitempty"`
	PriceMinimum float64 `json:"price_minimum,omitempty"`
}

type Coverage struct{
	ZipCodeInitial string `json:"zip_code_initial,omitempty"`
	ZipCodeFinal string `json:"zip_code_final,omitempty"`
}


type Promotion struct{
	Id string `json:"id,omitempty"`
	Off float64 `json:"off,omitempty"`
	Name string `json:"name,omitempty"`
	PromotionCode string `json:"promotion_code,omitempty"`
	PromotionCoverages []Coverage `json:"coverages,omitempty"`
	StartDate int64 `json:"initial_at,omitempty"`
	EndDate int64 `json:"end_at,omitempty"`
	StartHour string `json:"initial_hour,omitempty"`
	EndHour string `json:"end_hour,omitempty"`
}

type MoreUser struct{
	Modality string `json:"modality,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type Message struct{
	Message string `json:"message,omitempty"`
}