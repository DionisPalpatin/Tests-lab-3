package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	myerrors "github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/myerrors"
	"github.com/DionisPalpatin/Tests-lab-3/tree/main/backend/internal/models"
)

type IUserRepository interface {
	GetUserByID(id int) (*models.User, *myerrors.MyError)
	GetUserByLogin(login string) (*models.User, *myerrors.MyError)
	GetAllUsers() ([]*models.User, *myerrors.MyError)
	AddUser(user *models.User) *myerrors.MyError
	DeleteUser(id int) *myerrors.MyError
	UpdateUser(user *models.User) *myerrors.MyError

	GetAllUsersData() ([]*models.User, *myerrors.MyError)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func deferTransaction(err error, tx *sql.Tx) {
	if err != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, *myerrors.MyError) {
	if id < 0 {
		resState := myerrors.CreateError(myerrors.NoSuchUser, "GetUserByID", "data_access")
		return nil, resState
	}

	var user models.User

	query := fmt.Sprintf(getUserByIDQuery, "test_lab3")

	err := ur.db.QueryRowContext(context.Background(), query, id).Scan(
		&user.Id,
		&user.Fio,
		&user.RegistrationDate,
		&user.Login,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		var resState *myerrors.MyError

		if errors.Is(err, sql.ErrNoRows) {
			resState = myerrors.CreateError(myerrors.NoSuchUser, "GetUserByID", "data_access")
		} else {
			resState = myerrors.CreateError(myerrors.DatabaseError, "GetUserByID", "data_access")
		}

		return nil, resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "GetUserByID", "data_access")
	return &user, resState
}

func (ur *UserRepository) GetUserByLogin(loginOrFio string) (*models.User, *myerrors.MyError) {
	if loginOrFio == "" {
		resState := myerrors.CreateError(myerrors.NoSuchUser, "GetUserByLogin", "data_access")
		return nil, resState
	}

	var user models.User

	query := fmt.Sprintf(getUserByStringQuery, "test_lab3")
	
	err := ur.db.QueryRowContext(context.Background(), query, loginOrFio).Scan(
		&user.Id,
		&user.Fio,
		&user.RegistrationDate,
		&user.Login,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		var resState *myerrors.MyError

		if errors.Is(err, sql.ErrNoRows) {
			resState = myerrors.CreateError(myerrors.NoSuchUser, "GetUserByLogin", "data_access")
		} else {
			resState = myerrors.CreateError(myerrors.NoSuchUser, "GetUserByLogin", "data_access")
		}

		return nil, resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "GetUserByLogin", "data_access")
	return &user, resState
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, *myerrors.MyError) {
	query := fmt.Sprintf(getAllUsersQuery, "test_lab3")

	rows, err := ur.db.QueryContext(context.Background(), query)
	defer rows.Close()

	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsers", "data_access")
		return nil, resState
	}

	var users []*models.User
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
			resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsers", "data_access")
			return nil, resState
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsers", "data_access")
		return nil, resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "GetAllUsers", "data_access")
	return users, resState
}

func (ur *UserRepository) AddUser(user *models.User) *myerrors.MyError {
	if user == nil {
		resState := myerrors.CreateError(myerrors.OperationError, "AddUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(addUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return myerrors.CreateError(myerrors.DatabaseError, "AddUser", "data_access")
	}
	defer deferTransaction(err, tx)

	_, err = tx.ExecContext(context.Background(), query,
		user.Fio,
		user.RegistrationDate,
		user.Login,
		user.Password,
		user.Role,
	)

	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "AddUser", "data_access")
		return resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "AddUser", "data_access")
	return resState
}

func (ur *UserRepository) DeleteUser(id int) *myerrors.MyError {
	if id < 0 {
		resState := myerrors.CreateError(myerrors.OperationError, "DeleteUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(deleteUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return myerrors.CreateError(myerrors.DatabaseError, "DeleteUser", "data_access")
	}
	defer deferTransaction(err, tx)

	_, err = tx.ExecContext(context.Background(), query, id)
	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "DeleteUser", "data_access")
		return resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "DeleteUser", "data_access")
	return resState
}

func (ur *UserRepository) UpdateUser(user *models.User) *myerrors.MyError {
	if user == nil {
		resState := myerrors.CreateError(myerrors.OperationError, "UpdateUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(updateUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return myerrors.CreateError(myerrors.DatabaseError, "UpdateUser", "data_access")
	}
	defer deferTransaction(err, tx)

	_, err = tx.ExecContext(context.Background(), query,
		user.Fio,
		user.RegistrationDate,
		user.Login,
		user.Password,
		user.Role,
		user.Id,
	)

	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "UpdateUser", "data_access")
		return resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "UpdateUser", "data_access")
	return resState
}


func (ur *UserRepository) GetAllUsersData() ([]*models.User, *myerrors.MyError) {
	query := fmt.Sprintf(getAllUsersQuery, "test_lab3")

	rows, err := ur.db.QueryContext(context.Background(), query)
	defer rows.Close()

	if err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsersData", "data_access")
		return nil, resState
	}

	var users []*models.User
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

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		resState := myerrors.CreateError(myerrors.DatabaseError, "GetAllUsersData", "data_access")
		return nil, resState
	}

	resState := myerrors.CreateError(myerrors.Ok, "GetAllUsersData", "data_access")
	return users, resState
}
