package model

type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
	State State `json:"state"`
}

type State struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Initials string `json:"initials"`
}


