package repositories

import (
	"database/sql"
	"fmt"
	"shortener/src/entity"
)

type UserRepository interface {
	Insert(user *entity.Users) (*entity.Users, error)
}

type UserRepositoryDb struct {
	Db *sql.DB
}

func (repo *UserRepositoryDb) Insert(user entity.Users) (int64, error) {
	defer repo.Db.Close()

	sql_statement := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3);"
	userDb, err := repo.Db.Exec(sql_statement, user.Username, user.Email, user.Password)

	if err != nil {
		fmt.Println("err", err)
		return 0, nil
	}
	id, _ := userDb.LastInsertId()

	return id, nil
}