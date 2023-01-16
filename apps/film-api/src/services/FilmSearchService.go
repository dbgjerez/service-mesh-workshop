package services

import (
	"film-api/domain/dto"
	"film-api/domain/repositories"
)

type FilmService struct {
	fClient *repositories.FilmRepository
	uClient *repositories.UserRepository
}

func NewFilmService(
	fClient *repositories.FilmRepository,
	uClient *repositories.UserRepository) *FilmService {
	return &FilmService{
		fClient: fClient,
		uClient: uClient}
}

func (fService *FilmService) FindFilmsByUser(username string) ([]dto.Film, error) {
	user, err := fService.uClient.GetUser(username)
	if err != nil {
		return nil, err
	}
	premium := false
	if user.Subscription == "premium" {
		premium = true
	}
	films, err := fService.fClient.GetFilms(premium)
	return films, err
}
