package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/models"
	myerrors "github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/myerrors"
)

type IUserRepository interface {
	GetAllUsersData() ([]models.User, *myerrors.MyError)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	dbx := sqlx.NewDb(db, "pgx")
	return &UserRepository{db: dbx}
}

func (ur *UserRepository) GetAllUsersData() ([]models.User, *myerrors.MyError) {
	query := fmt.Sprintf(getAllUsersQuery, "test_lab3")

	users := make([]models.User, 0)

	rows, err := ur.db.QueryxContext(context.Background(), query)
	
	if errors.Is(err, sql.ErrNoRows) {
		resState := myerrors.CreateError(myerrors.Ok, "GetAllUsersData", "data_access")
		return users, resState
	}
	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsersData", "data_access")
		return nil, resState
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&user.Fio,
			&user.RegistrationDate,
			&user.Login,
			&user.Password,
			&user.Role,
		)

		if err != nil {
			resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsersData", "data_access")
			return nil, resState
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsersData", "data_access")
		return nil, resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "GetAllUsersData", "data_access")
	return users, resState
}
