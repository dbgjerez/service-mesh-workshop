package model

type Film struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Duration int8   `json:"duration"`
	Premium  bool   `json:"premium"`
}
