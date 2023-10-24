package middleware

import "github.com/labstack/echo"

type goMiddleware struct {
}

type responseError struct {
	Message string `json:"message"`
}

func (m *goMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}
