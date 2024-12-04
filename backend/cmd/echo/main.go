package main

import (
	"net/http"

	myerrors "github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/myerrors"
	"github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/cmd/database"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("echo"))

	userRepository := database.Connect()
	if userRepository == nil {
		panic("cant init repo")
	}

	e.GET("/metrics", echoprometheus.NewHandler())
	
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/users", func(c echo.Context) error {
		users, err := userRepository.GetAllUsersData()
		if err != nil && err.ErrNum != myerrors.Ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, users)
	})

	e.Logger.Fatal(e.Start(":8081"))
}