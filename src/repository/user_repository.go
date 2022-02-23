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
	sql_statement := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3);"
	userDb, err := repo.Db.Exec(sql_statement, user.Username, user.Email, user.Password)

	if err != nil {
		fmt.Println("err", err)
		return 0, nil
	}
	id, _ := userDb.LastInsertId()

	return id, nil
}

func (repo *UserRepositoryDb) FindByEmail(email string) (entity.Users, error) {
	sql_statement := "SELECT id, username, email, password FROM users where email = $1;"
	userDb, err := repo.Db.Query(sql_statement, email)

	var user entity.Users

	if userDb.Next() {
		if err = userDb.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return entity.Users{
				Email:    "",
				Password: "",
			}, err
		}
	}

	if err != nil {
		fmt.Println("err", err)
		return entity.Users{}, nil
	}

	return user, nil
}
