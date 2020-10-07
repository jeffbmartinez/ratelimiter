package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Turns on "debug mode". For development only
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(CustomMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	a := e.Group("/api")
	// a.Use(ratelimiter-middleware)
	a.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/api/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	apiKey := c.QueryParam("apikey")
	return c.String(http.StatusOK, fmt.Sprintf("User ID: %v\nAPI Key: %v\n", id, apiKey))
}

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		delayHeaderValue := c.Request().Header.Get("X-Add-Delay")
		if delayHeaderValue != "" {
			delayDuration, err := time.ParseDuration(delayHeaderValue)

			if err == nil {
				fmt.Printf("Sleeping for %v\n", delayHeaderValue)
				time.Sleep(delayDuration)
			}
		}

		return next(c)
	}
}
