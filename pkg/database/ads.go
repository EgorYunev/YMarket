package database

import (
	"database/sql"

	"github.com/EgorYunev/YMarket/pkg/models"
)

type AdModel struct {
	db *sql.DB
}

func (m *AdModel) Insert() (int, error) {
	stmt := `INSERT INTO ads (name, description, userId)
			VALUES ($S, $S, $N)`

	res, err := m.db.Exec(stmt)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err

	}

	return int(id), nil
}

func (m *AdModel) GetAllByUserId() ([]models.Ad, error) {
	stmt := `SELECT * FROM ads
			WHERE user_id = $N`
	row, err := m.db.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	ads := []models.Ad{}
	for row.Next() {
		ad := &models.Ad{}
		row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.UserId)
		ads = append(ads, *ad)
	}

	if row.Err() != nil {
		return nil, err
	}

	return ads, nil

}

func (m *AdModel) GetById() (models.Ad, error) {
	stmt := `SELECT * FROM ads WHERE id = $N`

	row := m.db.QueryRow(stmt)

	res := models.Ad{}
	err := row.Scan(&res.Id, &res.Name, &res.Description)

	if err != nil {
		return res, err
	}

	if row.Err() != nil {
		return res, row.Err()
	}

	return res, nil
}
