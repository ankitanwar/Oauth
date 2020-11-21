package services

import (
	accesstoken "github.com/ankitanwar/OAuth/domain/accessToken"
	"github.com/ankitanwar/OAuth/repository/db"
	"github.com/ankitanwar/OAuth/utils/errors"
)

//Service : Interface for services
type Service interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestError)
	Create(accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError)
	UpdateExperationTime(accesstoken.AccessToken) *errors.RestError
}

type service struct {
	repository db.Repository
}

//NewService : it will return the pointer to the Service interface
func NewService(repo db.Repository) Service {
	return &service{
		repository: repo,
	}
}

//GetById : To get the user buy the given id
func (s *service) GetByID(id string) (*accesstoken.AccessToken, *errors.RestError) {
	//id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequest("Invalid access token id")
	}
	token, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return token, nil

}

//Create : To create the user by id
func (s *service) Create(at accesstoken.AccessToken) (*accesstoken.AccessToken, *errors.RestError) {
	err := at.Validate()
	if err != nil {
		return nil, err
	}
	token, err := s.repository.Create(at)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *service) UpdateExperationTime(at accesstoken.AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExperationTime(at)
}
