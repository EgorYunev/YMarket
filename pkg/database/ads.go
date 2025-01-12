package database

import (
	"database/sql"

	"github.com/EgorYunev/YMarket/pkg/models"
)

type AdModel struct {
	DB *sql.DB
}

func (m *AdModel) Insert() (int, error) {
	stmt := `INSERT INTO ads (name, description, user_id, price)
			VALUES ($S, $S, $N, &N)`

	res, err := m.DB.Exec(stmt)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err

	}

	return int(id), nil
}

func (m *AdModel) GetAllByUserId() ([]*models.Ad, error) {
	stmt := `SELECT * FROM ads
			WHERE user_id = $N`
	row, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	ads := []*models.Ad{}
	for row.Next() {
		ad := &models.Ad{}
		row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.UserId)
		ads = append(ads, ad)
	}

	if row.Err() != nil {
		return nil, err
	}

	return ads, nil

}

func (m *AdModel) GetById() (models.Ad, error) {
	stmt := `SELECT * FROM ads WHERE id = $N`

	row := m.DB.QueryRow(stmt)

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

func (m *AdModel) GetLastest() ([]*models.Ad, error) {
	stmt := `SELECT * FROM ads`

	row, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	ads := []*models.Ad{}

	i := 0
	for row.Next() {
		i++
		ad := &models.Ad{}

		err = row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price)
		if err != nil {
			return nil, err
		}

		ads = append(ads, ad)
		if i >= 10 {
			break
		}
	}

	if row.Err() != nil {
		return nil, row.Err()
	}

	return ads, nil
}

func (m *AdModel) GetAdsFiltered(name string) ([]*models.Ad, error) {
	stmt := `SELECT * FROM ads
			WHERE "name" = $1`
	ads := []*models.Ad{}

	row, err := m.DB.Query(stmt, name)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		ad := &models.Ad{}

		row.Scan(&ad.Id, &ad.Name, &ad.Description, &ad.Price)
		ads = append(ads, ad)
	}

	if row.Err() != nil {
		return nil, row.Err()
	}

	return ads, nil
}
