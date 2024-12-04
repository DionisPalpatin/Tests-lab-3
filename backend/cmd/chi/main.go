package main

import (
	"encoding/json"
	"net/http"

	bl "github.com/DionisPalpatin/ppo-and-db/tree/master/application/internal/business_logic"
	"github.com/DionisPalpatin/ppo-and-db/tree/master/application/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userRepo := database.Connect()
	if userRepo == nil {
		panic("cant init repo")
	}

	// /metrics endpoint
	router.Handle("/metrics", promhttp.Handler())

	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := userRepo.GetAllUsersData()
		if err != nil && err.ErrNum != bl.Ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	})

	http.ListenAndServe(":8081", router)
}
