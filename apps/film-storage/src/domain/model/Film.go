package model

type Film struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Premium  bool   `json:"premium"`
}
