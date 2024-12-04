package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	repos "github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/repos"
)

func Connect() repos.IUserRepository {
	connString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), "5432")

	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	for {
		err := db.Ping()
		if err == nil {
			fmt.Println("Соединение с БД установлено!")
			break
		} else {
			fmt.Println("Не получилось соединиться с БД, еще одна попытка...")
			time.Sleep(5 * time.Second) // Пауза между попытками
		}
	}

	return repos.NewUserRepository(db)
}
