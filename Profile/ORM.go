package Profile

import (
	"context"
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
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	if role != 0 {
		return "client", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id)
	if err != nil {
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
		}
	}
	if role != 0 {
		return "host", nil
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id)
	if err != nil {
		return "", err
	}
	for row.Next() {
		err = row.Scan(&role)
		if err != nil {
			return "", err
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
		return Profile{}, err
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, err
		}
	}
	return profile, err
}

func (db *Wrapper) GetProfileClient(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, err
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, err
		}
	}

	row, err = db.Conn.Query(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id)
	if err != nil {
		return Profile{}, err
	}

	for row.Next() {
		err = row.Scan(&profile.Birthday)
		if err != nil {
			panic(err)
			return Profile{}, err
		}
	}

	return profile, err
}

func (db *Wrapper) GetProfileCourier(id int) (Profile, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id)
	if err != nil {
		return Profile{}, err
	}

	var profile = Profile{}
	for row.Next() {
		err = row.Scan(&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
		if err != nil {
			return Profile{}, err
		}
	}
	return profile, err
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
