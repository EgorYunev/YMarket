package database

import (
	"database/sql"

	"github.com/EgorYunev/YMarket/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, password string) error {
	stmt := `INSERT INTO users (name, password)
			VALUES($1, $2)`
	_, err := m.DB.Exec(stmt, name, password)

	if err != nil {
		return err
	}

	return nil

}

func (m *UserModel) GetById(id int) (*models.User, error) {
	stmt := `SELECT id, name, password,
			a.id, a.name, a.description, a.price
			FROM users
			LEFT JOIN ads a ON id = a.user_id
			WHERE id = %n`
	row := m.DB.QueryRow(stmt, id)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Ads)

	if err != nil {
		return nil, err
	}

	return user, nil
}
