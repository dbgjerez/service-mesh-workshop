package repositories

import (
	"encoding/json"
	"errors"
	"film-api/domain/dto"
	"film-api/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	USER_SERVICE_URL string = "USER_SERVICE_URL"
)

type UserRepository struct {
	client *http.Client
	url    *string
}

func NewUserRepository() *UserRepository {
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	url := utils.GetEnv(USER_SERVICE_URL, "")
	if url == "" {
		log.Fatalf("env var USER_SERVICE_URL not defined")
	}
	return &UserRepository{
		client: &client,
		url:    &url,
	}
}

func (dao *UserRepository) GetUser(username string) (dto.User, error) {
	res, err := dao.client.Get(*dao.url + username)
	var user dto.User
	if err != nil || res.StatusCode != 200 {
		if err != nil {
			return user, err
		} else if res.StatusCode == 404 {
			log.Printf("User %s not found", username)
			return user, errors.New(fmt.Sprintf("User %s not found", username))
		} else {
			return user, errors.New(fmt.Sprintf("Error %d", res.StatusCode))
		}
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
