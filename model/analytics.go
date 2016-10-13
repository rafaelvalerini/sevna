package model

import(
	"time"
)


type Analytics struct{
	Start string  `json:"start"`
	End string  `json:"end"`
	Player Player `json:"player"` 
	Device Device `json:"player"`
}

type AnalyticsDevice struct{
	IdDevice string `json:"id"`
	OperationSystem string `json:"operation_system"`
	OperationSystemVersion string `json:"operation_system_version"`
	Device string `json:"device"`
	TypeConnection string `json:"type_connection"`
}

type AnalyticsPlayer struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Modality Modality `json:"modality,omitempty"`
	Price string `json:"price,omitempty"`
}

type AnalyticsModality struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Promotion Promotion `json:"promotion,omitempty"`
}


type AnalyticsPromotion struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ExibitionName string `json:"exibition_name,omitempty"`
}

type ResultAnalytics struct{
	DateTime time.Time `json:"date,omitempty"`
    StartAddress string `json:"start_address,omitempty"`
    EndAddress string `json:"end_address,omitempty"`
    Player  string `json:"player,omitempty"`
    Modality  string `json:"modality,omitempty"`
    Value  string `json:"value,omitempty"`
    OperationSystem  string `json:"operation_system,omitempty"`
    OperationSystemVersion string `json:"operation_system_version,omitempty"`
    TypeConnection string `json:"type_connection,omitempty"`
    StartCity string `json:"start_city,omitempty"`
    StartState string `json:"start_state,omitempty"`
    EndCity string `json:"end_city,omitempty"`
    EndState string `json:"end_state,omitempty"`
}