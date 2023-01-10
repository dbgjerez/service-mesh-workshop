package model

import (
	"errors"
)

type FilmRepository struct {
	films []Film
	store *Store
}

func NewFilmRepository(store *Store) (*FilmRepository, error) {
	films, err := store.ReadStore()
	if err != nil {
		return nil, err
	}
	return &FilmRepository{
			films: films,
			store: store},
		nil
}

func (dao *FilmRepository) HealthCheck() error {
	if !dao.store.Exists() {
		return errors.New("Database not found!")
	}
	return nil
}

func (dao *FilmRepository) FindById(idFilm int) *Film {
	films := dao.films
	for _, f := range films {
		if f.Id == idFilm {
			return &f
		}
	}
	return nil
}

func (dao *FilmRepository) GetAll() []Film {
	return dao.films
}

func (dao *FilmRepository) GetAllByType(premium bool) []Film {
	var res []Film
	if !premium {
		res = []Film{}
		for _, f := range dao.GetAll() {
			if !f.Premium {
				res = append(res, f)
			}
		}
	} else {
		res = dao.GetAll()
	}
	return res
}
