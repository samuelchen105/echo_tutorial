package api

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) error {
	SetValidator(e)
	SetRenderer(e)

	SetBindHandler(e.Group("/bind"))
	SetContextHandler(e.Group("/context"))
	SetCookieHandler(e.Group("/cookie"))
	SetRequestHandler(e.Group("/request"))
	SetResponseHandler(e.Group("/response"))

	err := SetCasbinHandler(e.Group("/casbin"))
	if err != nil {
		return err
	}

	SetCorsHandler(e.Group("/cors"))
	SetCsrfHandler(e.Group("/csrf"))
	SetJwtHandler(e.Group("/jwt"))

	return nil
}
