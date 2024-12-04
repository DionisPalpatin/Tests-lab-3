package main

import (
	"net/http"

	myerrors "github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/myerrors"
	"github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/cmd/database"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := gin.Default()

	userRepository := database.Connect()
	if userRepository == nil {
		panic("cant init repo")
	}

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/users", func(c *gin.Context) {
		users, err := userRepository.GetAllUsersData()
		if err != nil && err.ErrNum != myerrors.Ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	router.Run(":8081")
}