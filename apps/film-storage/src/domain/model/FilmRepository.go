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

func (dao *FilmRepository) FindById(idFilm int) (*Film, error) {
	films, err := dao.GetAll()
	if err != nil {
		return nil, err
	}
	for _, f := range films {
		if f.Id == idFilm {
			return &f, nil
		}
	}
	return nil, nil
}

func (dao *FilmRepository) GetAll() ([]Film, error) {
	films, err := dao.store.ReadStore()
	if err != nil {
		return []Film{}, err
	}
	return films, nil
}
