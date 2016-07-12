package model

type SellerStore struct {
	Id          int64         `json:"_id" bson:"id"`
	Name        string        `json:"name" bson:"name"`
	Lat         float64       `json:"lat" bson:"lat"`
	Lng         float64       `json:"lng" bson:"lng"`
	Address     string        `json:"address" bson:"address"`
	Phone       string        `json:"phone" bson:"phone"`
	Cards       Cards         `json:"cards" bson:"cards"`
	Description string        `json:"description" bson:"description"`
	OfficeHours []OfficeHours `json:"officeHours" bson:"officeHours"`
	Loc         Location      `json:"location" bson:"location"`
	IdSeller    int64         `json:"id_seller" bson:"id_seller"`
}

type Cards struct {
	IsCheckedVisa       bool `json:"is_checked_visa" bson:"is_checked_visa"`
	IsCheckedMaster     bool `json:"is_checked_master" bson:"is_checked_master"`
	IsCheckedAmerican   bool `json:"is_checked_american" bson:"is_checked_american"`
	IsCheckedElo        bool `json:"is_checked_elo" bson:"is_checked_elo"`
	IsCheckedGoodcard   bool `json:"is_checked_goodcard" bson:"is_checked_goodcard"`
	IsCheckedHipercard  bool `json:"is_checked_hipercard" bson:"is_checked_hipercard"`
	IsCheckedMaestro    bool `json:"is_checked_maestro" bson:"is_checked_maestro"`
	IsCheckedSodexo     bool `json:"is_checked_sodexo" bson:"is_checked_sodexo"`
	IsCheckedVR         bool `json:"is_checked_vr" bson:"is_checked_vr"`
	IsCheckedAlelo      bool `json:"is_checked_alelo" bson:"is_checked_alelo"`
	IsCheckedDinersclub bool `json:"is_checked_dinersclub" bson:"is_checked_dinersclub"`
}

type OfficeHours struct {
	Seg         bool   `json:"seg" bson:"seg"`
	Ter         bool   `json:"ter" bson:"ter"`
	Qua         bool   `json:"qua" bson:"id"`
	Qui         bool   `json:"qui" bson:"id"`
	Sex         bool   `json:"sex" bson:"id"`
	Sab         bool   `json:"sab" bson:"id"`
	Dom         bool   `json:"dom" bson:"id"`
	InitialHour string `json:"initial_hour" bson:"initial_hour"`
	FinalHour   string `json:"final_hour" bson:"final_hour"`
}

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
