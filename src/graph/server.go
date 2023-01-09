package graph

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-playground/validator/v10"

	"github.com/liac-inc/gqlgen-template/src/graph/generated"
	"github.com/liac-inc/gqlgen-template/src/graph/resolver"
	"github.com/liac-inc/gqlgen-template/src/middleware"
	"github.com/liac-inc/gqlgen-template/src/registry"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func ServerInit(db *sql.DB) {
	logger := loggerInit()
	resolvers := registry.InitializeResolver(db)

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Cors())
	e.Use(middleware.Logger(logger))
	e.Use(middleware.EchoCtxToCtx())

	e.GET("/", playgroundHandler())
	e.POST("/query", graphqlHandler(resolvers))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", os.Getenv("API_PORT"))

	e.Logger.Fatal(e.Start(":4000"))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func loggerInit() *zap.Logger {
	var logger *zap.Logger
	var err error

	env := os.Getenv("ENV")
	switch env {
	case "local", "dev":
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	case "prod":
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}

	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	return logger
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func graphqlHandler(resolver *resolver.Resolver) echo.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}
