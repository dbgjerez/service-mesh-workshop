package service

import (
	"film-storage/domain/dto"
	"film-storage/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilmCreateDTOToFilm(dto dto.FilmCreateDTO) model.Film {
	return model.Film{
		Id:       primitive.NewObjectID(),
		Title:    dto.Title,
		Duration: int8(dto.Duration),
		Premium:  dto.Premium,
	}
}
