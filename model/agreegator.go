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
	PopupMultiplier string `json:"popup_multiplier,omitempty"`
	Active int `json:"active,omitempty"`
	Token string `json:"token,omitempty"`
	AlertMessage string `json:"alert_message,omitempty"`
}

type Modality struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Promotion Promotion `json:"promotion,omitempty"`
	Promotions []Promotion `json:"promotions,omitempty"`
	PriceKm float64 `json:"price_km,omitempty"`
	TimeKm int `json:"time_km,omitempty"`
	ModalityCoverage []Coverage `json:"coverages,omitempty"`
	PriceTime float64 `json:"price_time,omitempty"`
	PriceBase float64 `json:"price_base,omitempty"`
	PriceMinimum float64 `json:"price_minimum,omitempty"`
	Active int `json:"active,omitempty"`
	EditValues int `json:"edit_values,omitempty"`
	KeyApi string `json:"key_api,omitempty"`
}

type Coverage struct{
	ZipCodeInitial string `json:"zip_code_initial,omitempty"`
	ZipCodeFinal string `json:"zip_code_final,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Available struct{
	Id int `json:"id,omitempty"`
	Monday int `json:"monday,omitempty"`
	Tuesday int `json:"tuesday,omitempty"`
	Wednesday int `json:"wednesday,omitempty"`
	Thursday int `json:"thursday,omitempty"`
	Friday int `json:"friday,omitempty"`
	Saturday int `json:"saturday,omitempty"`
	Sunday int `json:"sunday,omitempty"`
	StartHour string `json:"start_hour,omitempty"`
	EndHour string `json:"end_hour,omitempty"`
	IdPromotion string `json:"idPromotion,omitempty"`
}

type Promotion struct{
	Id int `json:"id,omitempty"`
	Off float64 `json:"off,omitempty"`
	LimitOff float64 `json:"limit_off,omitempty"`
	Name string `json:"name,omitempty"`
	ExibitionName string `json:"exibition_name,omitempty"`
	Description string `json:"description,omitempty"`
	PromotionCode string `json:"promotion_code,omitempty"`
	PromotionAvailable []Available `json:"availables,omitempty"`
	PromotionCoverages []Coverage `json:"coverages,omitempty"`	
	StartDate string `json:"initial_at,omitempty"`
	EndDate string `json:"end_at,omitempty"`
	Modality string `json:"modality_name,omitempty"`
	NewModality int `json:"new_modality,omitempty"`
	Modalities []ModalitySimple `json:"modalities,omitempty"`
	Active int `json:"active,omitempty"`
}

type ModalitySimple struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Player int `json:"player,omitempty"`
}

type MoreUser struct{
	Modality string `json:"modality,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type Message struct{
	Message string `json:"message,omitempty"`
}