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
	Active int `json:"active,omitempty"`
	Token string `json:"token,omitempty"`
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
	Active int `json:"active,omitempty"`
	EditValues int `json:"edit_values,omitempty"`
}

type Coverage struct{
	ZipCodeInitial string `json:"zip_code_initial,omitempty"`
	ZipCodeFinal string `json:"zip_code_final,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Available struct{
	Id int `json:"id,omitempty"`
	Monday int `json:"Monday,omitempty"`
	Tuesday int `json:"Tuesday,omitempty"`
	Wednesday int `json:"Wednesday,omitempty"`
	Thursday int `json:"Thursday,omitempty"`
	Friday int `json:"Friday,omitempty"`
	Saturday int `json:"Saturday,omitempty"`
	Sunday int `json:"Sunday,omitempty"`
	StartHour string `json:"Start_hour,omitempty"`
	EndHour string `json:"End_hour,omitempty"`
	IdPromotion string `json:"IdPromotion,omitempty"`
}

type Promotion struct{
	Id string `json:"id,omitempty"`
	Off float64 `json:"off,omitempty"`
	Name string `json:"name,omitempty"`
	PromotionCode string `json:"promotion_code,omitempty"`
	PromotionAvailable []Available `json:"availables,omitempty"`
	PromotionCoverages []Coverage `json:"coverages,omitempty"`	
	StartDate string `json:"initial_at,omitempty"`
	EndDate string `json:"end_at,omitempty"`
	Modality string `json:"modality_name,omitempty`
	NewModality int `json:"new_modality,omitempty"`
}

type MoreUser struct{
	Modality string `json:"modality,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type Message struct{
	Message string `json:"message,omitempty"`
}