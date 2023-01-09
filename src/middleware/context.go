package middleware

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
)

type echoCtxKey struct{}

func EchoCtxToCtx() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), echoCtxKey{}, c)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)

			return next(c)
		}
	}
}

func GetEchoCtxFromCtx(ctx context.Context) (echo.Context, error) {
	echoCtx := ctx.Value(echoCtxKey{})
	if echoCtx == nil {
		return nil, errors.New("could not retrieve echo.Context")
	}

	ec, ok := echoCtx.(echo.Context)
	if !ok {
		return nil, errors.New("echo.Context has wrong type")
	}

	return ec, nil
}
