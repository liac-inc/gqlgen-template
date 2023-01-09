package service

import (
	"github.com/liac-inc/gqlgen-template/src/graph/model"
	"github.com/liac-inc/gqlgen-template/src/repository"
)

type IUserService interface {
	FindAllUsers() ([]*model.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) FindAllUsers() ([]*model.User, error) {
	return u.repo.FindAllUsers()
}
