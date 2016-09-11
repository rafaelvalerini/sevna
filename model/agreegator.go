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
}

type Player struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Modality Modality `json:"modality"`
	WaitingTime int `json:"waiting_time"`
	Price string `json:"price"`
	Uuid string `json:"uuid"`
	Multiplier float64 `json:"multiplier"`
}

type Modality struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Promotion Promotion `json:"promotion"`
	PriceKm float64 `json:"price_km"`
	TimeKm int `json:"time_km"`
}

type Promotion struct{
	Id string `json:"id"`
	Off float64 `json:"off"`
	Name string `json:"name"`
	PromotionCode string `json:"promotion_code"`
}