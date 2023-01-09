package middleware

import (
	"github.com/labstack/echo/v4"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) echo.MiddlewareFunc {
	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			logger.Info("mw:logger",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.Error(v.Error),
			)
			return nil
		},
	})
}
