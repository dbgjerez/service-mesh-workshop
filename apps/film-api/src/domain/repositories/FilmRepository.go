package repositories

import (
	"encoding/json"
	"film-api/domain/dto"
	"film-api/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	FILM_SERVICE_URL string = "FILM_SERVICE_URL"
)

type FilmRepository struct {
	client *http.Client
	url    *string
}

func NewFilmRepository() *FilmRepository {
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	url := utils.GetEnv(FILM_SERVICE_URL, "")
	if url == "" {
		log.Fatalf("env var FILM_SERVICE_URL not defined")
	}
	return &FilmRepository{
		client: &client,
		url:    &url,
	}
}

func (dao *FilmRepository) GetFilms(premium bool) ([]dto.Film, error) {
	res, err := dao.client.Get(*dao.url + "?premium=" + strconv.FormatBool(premium))
	if err != nil {
		log.Printf("http error calling %s: %v", *dao.url, err)
	}
	defer res.Body.Close()
	var films []dto.Film
	if err := json.NewDecoder(res.Body).Decode(&films); err != nil {
		return nil, err
	}
	return films, nil
}
