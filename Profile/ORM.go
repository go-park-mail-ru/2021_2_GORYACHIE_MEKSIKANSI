package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	//mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Wrapper struct {
	Conn *pgxpool.Pool
}

func (db *Wrapper) getRoleById(id int) (string, error) {
	role := 0

	row, err := db.Conn.Query(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(errorsConst.ERRCLIENTQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(errorsConst.ERRCLIENTSCAN)
		}
	}
	if role != 0 {
		return "client", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(errorsConst.ERRHOSTQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(errorsConst.ERRHOSTSCAN)
		}
	}
	if role != 0 {
		return "host", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(errorsConst.ERRCORIERQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(errorsConst.ERRCORIERSCAN)
		}
	}
	if role != 0 {
		return "courier", nil
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILEHOSTQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return nil, errors.New(errorsConst.ERRGETPROFILEHOSTSCAN)
		}
	}
	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILECLIENTQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return nil, errors.New(errorsConst.ERRGETPROFILECLIENTSCAN)
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETBIRTHDAYQUERY)
	}

	for row.Next() {
		err = row.Scan(&profile.Birthday)
		if err != nil {
			return nil, errors.New(errorsConst.ERRGETBIRTHDAYSCAN)
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILECOURIERQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return nil, errors.New(errorsConst.ERRGETPROFILECOURIERSCAN)
		}
	}
	return &profile, nil
}
