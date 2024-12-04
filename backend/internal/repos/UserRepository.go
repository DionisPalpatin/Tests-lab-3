package data_access

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	bl "github.com/DionisPalpatin/ppo-and-db/tree/master/application/internal/business_logic"
	"github.com/DionisPalpatin/ppo-and-db/tree/master/application/internal/models"
)

type IUserRepository interface {
	GetUserByID(id int) (*models.User, *bl.MyError)
	GetUserByLogin(login string) (*models.User, *bl.MyError)
	GetAllUsers() ([]*models.User, *bl.MyError)
	AddUser(user *models.User) *bl.MyError
	DeleteUser(id int) *bl.MyError
	UpdateUser(user *models.User) *bl.MyError

	GetAllUsersData() ([]*models.User, *bl.MyError)
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

func (ur *UserRepository) GetUserByID(id int) (*models.User, *bl.MyError) {
	if id < 0 {
		resState := bl.CreateError(bl.NoSuchUser, "GetUserByID", "data_access")
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
		var resState *bl.MyError

		if errors.Is(err, sql.ErrNoRows) {
			resState = bl.CreateError(bl.NoSuchUser, "GetUserByID", "data_access")
		} else {
			resState = bl.CreateError(bl.DatabaseError, "GetUserByID", "data_access")
		}

		return nil, resState
	}

	resState := bl.CreateError(bl.Ok, "GetUserByID", "data_access")
	return &user, resState
}

func (ur *UserRepository) GetUserByLogin(loginOrFio string) (*models.User, *bl.MyError) {
	if loginOrFio == "" {
		resState := bl.CreateError(bl.NoSuchUser, "GetUserByLogin", "data_access")
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
		var resState *bl.MyError

		if errors.Is(err, sql.ErrNoRows) {
			resState = bl.CreateError(bl.NoSuchUser, "GetUserByLogin", "data_access")
		} else {
			resState = bl.CreateError(bl.NoSuchUser, "GetUserByLogin", "data_access")
		}

		return nil, resState
	}

	resState := bl.CreateError(bl.Ok, "GetUserByLogin", "data_access")
	return &user, resState
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, *bl.MyError) {
	query := fmt.Sprintf(getAllUsersQuery, "test_lab3")

	rows, err := ur.db.QueryContext(context.Background(), query)
	defer rows.Close()

	if err != nil {
		resState := bl.CreateError(bl.DatabaseError, "GetAllUsers", "data_access")
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
			resState := bl.CreateError(bl.DatabaseError, "GetAllUsers", "data_access")
			return nil, resState
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		resState := bl.CreateError(bl.DatabaseError, "GetAllUsers", "data_access")
		return nil, resState
	}

	resState := bl.CreateError(bl.Ok, "GetAllUsers", "data_access")
	return users, resState
}

func (ur *UserRepository) AddUser(user *models.User) *bl.MyError {
	if user == nil {
		resState := bl.CreateError(bl.OperationError, "AddUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(addUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return bl.CreateError(bl.DatabaseError, "AddUser", "data_access")
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
		resState := bl.CreateError(bl.DatabaseError, "AddUser", "data_access")
		return resState
	}

	resState := bl.CreateError(bl.Ok, "AddUser", "data_access")
	return resState
}

func (ur *UserRepository) DeleteUser(id int) *bl.MyError {
	if id < 0 {
		resState := bl.CreateError(bl.OperationError, "DeleteUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(deleteUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return bl.CreateError(bl.DatabaseError, "DeleteUser", "data_access")
	}
	defer deferTransaction(err, tx)

	_, err = tx.ExecContext(context.Background(), query, id)
	if err != nil {
		resState := bl.CreateError(bl.DatabaseError, "DeleteUser", "data_access")
		return resState
	}

	resState := bl.CreateError(bl.Ok, "DeleteUser", "data_access")
	return resState
}

func (ur *UserRepository) UpdateUser(user *models.User) *bl.MyError {
	if user == nil {
		resState := bl.CreateError(bl.OperationError, "UpdateUser", "data_access")
		return resState
	}

	query := fmt.Sprintf(updateUserQuery, "test_lab3")

	tx, err := ur.db.BeginTx(context.Background(), nil)
	if err != nil {
		return bl.CreateError(bl.DatabaseError, "UpdateUser", "data_access")
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
		resState := bl.CreateError(bl.DatabaseError, "UpdateUser", "data_access")
		return resState
	}

	resState := bl.CreateError(bl.Ok, "UpdateUser", "data_access")
	return resState
}


func (ur *UserRepository) GetAllUsersData() ([]*models.User, *bl.MyError) {
	query := fmt.Sprintf(getAllUsersQuery, "test_lab3")

	rows, err := ur.db.QueryContext(context.Background(), query)
	defer rows.Close()

	if err != nil {
		resState := bl.CreateError(bl.DatabaseError, "GetAllUsersData", "data_access")
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
			resState := bl.CreateError(bl.DatabaseError, "GetAllUsersData", "data_access")
			return nil, resState
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		resState := bl.CreateError(bl.DatabaseError, "GetAllUsersData", "data_access")
		return nil, resState
	}

	resState := bl.CreateError(bl.Ok, "GetAllUsersData", "data_access")
	return users, resState
}
