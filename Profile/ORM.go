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

	err := db.Conn.QueryRow(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", errors.New(errorsConst.ERRHOSTSCAN)
	}
	if role != 0 {
		return "host", nil
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", errors.New(errorsConst.ERRCLIENTSCAN)
	}
	if role != 0 {
		return "client", nil
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", errors.New(errorsConst.ERRCORIERSCAN)
	}
	if role != 0 {
		return "courier", nil
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*Profile, error) {
	var profile = Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
			&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILEHOSTSCAN)
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*Profile, error) {
	var profile = Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
			&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILECLIENTSCAN)
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&profile.Birthday)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETBIRTHDAYSCAN)
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*Profile, error) {
	var profile = Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
			&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, errors.New(errorsConst.ERRGETPROFILECOURIERSCAN)

	}
	return &profile, nil
}
