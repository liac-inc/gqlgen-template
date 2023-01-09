//go:build wireinject
// +build wireinject

package registry

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/liac-inc/gqlgen-template/src/graph/resolver"
	"github.com/liac-inc/gqlgen-template/src/repository"
	"github.com/liac-inc/gqlgen-template/src/service"
)

var userSet = wire.NewSet(service.NewUserService, repository.NewUserRepository, wire.Bind(new(service.IUserService), new(*service.UserService)), wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)))

func InitializeResolver(db *sql.DB) *resolver.Resolver {
	wire.Build(resolver.NewResolver, userSet)

	return &resolver.Resolver{}
}
