package repository

import (
	"database/sql"
	"fmt"
	"taskManager/internal/domain/user"
	"taskManager/internal/infrastruct/database"
	"taskManager/internal/infrastruct/database/entity"
	"taskManager/internal/infrastruct/mapper"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) user.IUserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) GetAll() ([]user.UserModel, error) {
	const (
		op    = "UserRepository.Create"
		query = "select * from User"
	)

	rows, err := u.db.Instance.Query(query)
	if err != nil {
		return nil, database.ReadFailedError(op, err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)

	var users []entity.UserEntity

	for rows.Next() {
		var user entity.UserEntity
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, database.ReadFailedError(op, err.Error())
		}
		users = append(users, user)
	}

	var userModels []user.UserModel

	for _, user := range users {
		userModels = append(userModels, *mapper.MapUserFromEntityToDomain(&user))
	}

	return userModels, nil
}

func (u UserRepository) Create(newUser user.UserModel) error {
	const (
		op    = "UserRepository.Create"
		query = "insert into User (name, email, password) values (?, ?, ?)"
	)

	res, err := u.db.Instance.Exec(query, newUser.Name, newUser.Email, newUser.Password)

	if err != nil {
		return database.CreateFailedError(op, err.Error())
	}

	if _, err := res.LastInsertId(); err != nil {
		return database.CreateFailedError(op, err.Error())
	}

	return nil
}

func (u UserRepository) GetUser(params user.GetUserParams) (*user.UserModel, error) {
	const (
		op    = "UserRepository.Create"
		query = "select * from User where id = ?"
	)

	if params.Id == nil {
		return nil, database.ReadFailedError(op, "id is required")
	}

	rows, err := u.db.Instance.Query(query, params.Id)

	if err != nil {
		return nil, database.ReadFailedError(op, err.Error())
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			panic(err)
		}
	}(rows)

	var user entity.UserEntity

	if !rows.Next() {
		return nil, database.NotFoundError(op)
	}

	if errScan := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); errScan != nil {
		return nil, database.ReadFailedError(op, errScan.Error())
	}

	return mapper.MapUserFromEntityToDomain(&user), nil
}

func (u UserRepository) Delete(id int) error {
	const (
		op    = "UserRepository.Create"
		query = "delete from User where id = ?"
	)

	res, err := u.db.Instance.Exec(query, id)

	if err != nil {
		return database.DeleteFailedError(op, err.Error())
	}

	if _, resErr := res.RowsAffected(); resErr != nil {
		return database.DeleteFailedError(op, resErr.Error())
	}

	return nil
}

func (u UserRepository) Update(user user.UserModel) error {
	//TODO implement me
	panic("implement me")
}
