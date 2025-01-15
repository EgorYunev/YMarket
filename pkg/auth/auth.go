package auth

import (
	"database/sql"
	"errors"

	"github.com/EgorYunev/YMarket/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	DB *sql.DB
}

var singKey = []byte("ymarketauth")

func (a *Auth) GetToken(u *models.User) (string, error) {
	stmt := `SELECT id, name, password FROM users
			WHERE name = $1`

	row := a.DB.QueryRow(stmt, u.Name)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Password)

	if err != nil {
		return "", err
	}

	if u.Name != user.Name || u.Password != user.Password {
		err = errors.New("Incorrect data")
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	strToken, err := token.SignedString(singKey)

	if err != nil {
		return "", err
	}

	a.DB.Exec("UPDATE users SET token = $1 WHERE id = $2", strToken, user.Id)

	return strToken, nil
}
