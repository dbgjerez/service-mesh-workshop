package dto

type FilmCreateDTO struct {
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Premium  bool   `json:"premium,omitempty"`
}
