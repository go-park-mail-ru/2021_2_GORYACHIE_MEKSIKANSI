package Profile

import (
	mid "2021_2_GORYACHIE_MEKSIKANSI/Middleware"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ERRCLIENTQUERY = "ERROR: check user on client query"
	ERRCLIENTSCAN = "ERROR: check user on client not scan"
	ERRHOSTQUERY = "ERROR: check user on host query"
	ERRHOSTSCAN = "ERROR: check user on client not scan"
	ERRCORIERQUERY = "ERROR: check user on host query"
	ERRCORIERSCAN = "ERROR: check user on client not scan"
	ERRGETPROFILEHOSTQUERY = "ERROR: profile host query"
	ERRGETPROFILEHOSTSCAN = "ERROR: profile host not scan"
	ERRGETPROFILECLIENTQUERY = "ERROR: profile client query"
	ERRGETPROFILECLIENTSCAN = "ERROR: profile client not scan"
	ERRGETBIRTHDAYQUERY = "ERROR: birthday query"
	ERRGETBIRTHDAYSCAN = "ERROR: birthday not scan"
	ERRGETPROFILECOURIERQUERY = "ERROR: profile courier query"
	ERRGETPROFILECOURIERSCAN = "ERROR: profile courier not scan"
	ERRADDCOOKIEQUERY = "ERROR: cookie query"
)

type Wrapper struct {
	Conn *pgxpool.Pool
}

func (db *Wrapper) getRoleById(id int) (string, error) {
	role := 0

	row, err := db.Conn.Query(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(ERRCLIENTQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(ERRCLIENTSCAN)
		}
	}
	if role != 0 {
		return "client", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(ERRHOSTQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(ERRHOSTSCAN)
		}
	}
	if role != 0 {
		return "host", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id)
	if err != nil {
		return "", errors.New(ERRCORIERQUERY)
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", errors.New(ERRCORIERSCAN)
		}
	}
	if role != 0 {
		return "courier", nil
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, errors.New(ERRGETPROFILEHOSTQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, errors.New(ERRGETPROFILEHOSTSCAN)
		}
	}
	return profile, err
}

func (db *Wrapper) GetProfileClient(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, errors.New(ERRGETPROFILECLIENTQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, errors.New(ERRGETPROFILECLIENTSCAN)
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id)
	if err != nil {
		return Profile{}, errors.New(ERRGETBIRTHDAYQUERY)
	}

	for row.Next() {
		err = row.Scan(&profile.Birthday)
		if err != nil {
			panic(err)
			return Profile{}, errors.New(ERRGETBIRTHDAYSCAN)
		}
	}

	return profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, errors.New(ERRGETPROFILECOURIERQUERY)
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, errors.New(ERRGETPROFILECOURIERSCAN)
		}
	}
	return profile, nil
}

func (db *Wrapper) AddCookie(cookie mid.Defense, id int) error {
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO cookie (client_id, session_id, date_life) VALUES ($1, $2, $3)",
		id, cookie.SessionId, cookie.DateLife)
	if err != nil {
		return errors.New(ERRADDCOOKIEQUERY)
	}

	return nil
}

//func (db *Wrapper) updateName(id int, name string) error {
//	_, err := db.Conn.Query(context.Background(),
//		"UPDATE general_user_info SET name = $1 WHERE id = $2", name, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (db *Wrapper) updateEmail(id int, email string) error {
//	_, err := db.Conn.Query(context.Background(),
//		"UPDATE general_user_info SET email = $1 WHERE id = $2", email, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (db *Wrapper) updatePassword(id int, password string) error {
//	_, err := db.Conn.Query(context.Background(),
//		"UPDATE general_user_info SET password = $1 WHERE id = $2", password, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (db *Wrapper) updateAdditionalInfo(id int, phone string) error {
//	_, err := db.Conn.Query(context.Background(),
//		"UPDATE general_user_info SET phone = $1 WHERE id = $2", phone, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (db *Wrapper) updateAvatar(id int, avatar string) error {
//	_, err := db.Conn.Query(context.Background(),
//		"UPDATE general_user_info SET avatar = $1 WHERE id = $2", avatar, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
