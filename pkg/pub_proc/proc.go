package pub_proc

import (
	"boilerplate_go/pkg/api/routers"

	"github.com/labstack/echo/v4"
)

func PublicRouter(e *echo.Echo) {
	routers.InitProductRouter(e)
}