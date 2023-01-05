package model

type Film struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Premium  bool   `json:"premium,omitempty"`
}
