package repository

import (
	"context"
	"database/sql"

	"github.com/liac-inc/gqlgen-template/src/db/query"
	"github.com/liac-inc/gqlgen-template/src/graph/model"
)

type IUserRepository interface {
	FindAllUsers() ([]*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindAllUsers() ([]*model.User, error) {
	ctx := context.Background()
	queries := query.New(u.db)

	users, err := queries.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	dto := []*model.User{}
	for _, user := range users {
		dto = append(dto, &model.User{
			ID:   user.ID.String(),
			Name: user.Name,
		})
	}

	return dto, nil
}
