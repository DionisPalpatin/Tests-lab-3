package repos

// Users queries
const (
	getUserByIDQuery = `
	SELECT * 
	FROM %s.users
    WHERE id = $1;
	`
	getUserByStringQuery = `
	SELECT * 
	FROM %s.users
    WHERE login = $1 OR fio = $1;
	`
	getAllUsersQuery = `
	SELECT * 
	FROM %s.users;
	`
	addUserQuery = `
	INSERT INTO %s.users (fio, registration_date, login, password, role) VALUES
	($1, $2, $3, $4, $5);
	`
	deleteUserQuery = `
	call %s.delete_user($1);
	`
	updateUserQuery = `
	UPDATE %s.users SET fio = $1, registration_date = $2, login = $3, password = $4, role = $5
	WHERE id = $6;
	`
)