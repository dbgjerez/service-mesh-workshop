package dto

type Film struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Duration int       `json:"duration"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	User    UserComment `json:"user"`
	Comment string      `json:"comment"`
}

type UserComment struct {
	IdUser   int32  `json:"idUser"`
	UserName string `json:"username"`
}
