package services

import (
	"film-api/domain/dto"
	"film-api/domain/repository"
)

type FilmService struct {
	fClient *repository.FilmRepository
	uClient *repository.UserRepository
	cClient *repository.CommentRepository
}

func NewFilmService(
	fClient *repository.FilmRepository,
	uClient *repository.UserRepository,
	cClient *repository.CommentRepository) *FilmService {
	return &FilmService{
		fClient: fClient,
		uClient: uClient,
		cClient: cClient,
	}
}

func (fService *FilmService) FindFilmsByUser(username string) ([]*dto.Film, error) {
	user, err := fService.uClient.GetUser(username)
	if err != nil {
		return nil, err
	}
	premium := isPremium(&user)
	films, err := fService.fClient.GetFilms(premium)
	var ids []int32
	for _, f := range films {
		ids = append(ids, int32(f.Id))
	}
	comments, err := fService.cClient.FindComments(ids)

	for _, c := range comments.Comments {
		for _, f := range films {
			if f.Id == int(c.IdObject) {
				f.Comments = append(f.Comments, dto.Comment{
					Comment: c.Comment,
					User: dto.UserComment{
						IdUser: c.IdUser,
					},
				})
			}
		}
	}
	return films, err
}

func isPremium(user *dto.User) bool {
	premium := false
	if user.Subscription == "premium" {
		premium = true
	}
	return premium
}
