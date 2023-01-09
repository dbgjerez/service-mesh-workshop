package model

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Subscription string `json:"subscription"`
}
