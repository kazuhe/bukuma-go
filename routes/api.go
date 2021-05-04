package routes

import (
	"github.com/kazuhe/bukuma-go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	// All endpoints defined within this block will be prefixed with "/api"
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}
