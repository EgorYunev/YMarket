package database

import (
	"database/sql"

	"github.com/EgorYunev/YMarket/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, password string) (int, error) {
	stmt := `INSERT INTO users (name, password)
			VALUES($S, $S)`
	res, err := m.DB.Exec(stmt, name, password)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *UserModel) GetById(id int) (*models.User, error) {
	stmt := `SELECT id, name, password,
			a.id, a.name, a.description, a.price
			FROM users
			LEFT JOIN ads a ON id = a.user_id
			WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Ads)

	if err != nil {
		return nil, err
	}

	return user, nil
}
