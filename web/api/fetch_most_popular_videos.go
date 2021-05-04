package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Return "Most Popular" with status code 200
		return c.JSON(fasthttp.StatusOK, "Most Popular")
	}
}
