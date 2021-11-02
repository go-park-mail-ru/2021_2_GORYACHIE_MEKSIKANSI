package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	prof "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"time"
)

type Wrapper struct {
	Conn Utils.ConnectionInterface
}

func (db *Wrapper) GetRoleById(id int) (string, error) {
	role := 0

	err := db.Conn.QueryRow(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", &errorsConst.Errors{
			Text: errorsConst.ErrHostScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "host", nil
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", &errorsConst.Errors{
			Text: errorsConst.ErrClientScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "client", nil
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && err.Error() != "no rows in result set" {
		return "", &errorsConst.Errors{
			Text: errorsConst.ErrCourierScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "courier", nil
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*prof.Profile, error) {
	var profile = prof.Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrGetProfileHostScan,
			Time: time.Now(),
		}
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*prof.Profile, error) {
	var profile = prof.Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrGetProfileClientScan,
			Time: time.Now(),
		}
	}

	err = db.Conn.QueryRow(context.Background(),
		"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&profile.Birthday)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrGetBirthdayScan,
			Time: time.Now(),
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*prof.Profile, error) {
	var profile = prof.Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.ErrGetProfileCourierScan,
			Time: time.Now(),
		}
	}
	return &profile, nil
}

func (db *Wrapper) UpdateName(id int, newName string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET name = $1 WHERE id = $2",
		newName, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdateName,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateEmail(id int, newEmail string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET email = $1 WHERE id = $2",
		newEmail, id)
	if err != nil {
		textError := err.Error()
		println(textError)
		if textError == "ERROR: duplicate key value violates unique constraint "+
			"\"general_user_info_email_key\" (SQLSTATE 23505)" {
			return &errorsConst.Errors{
				Text: errorsConst.ErrUpdateEmailRepeat,
				Time: time.Now(),
			}
		}
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdateEmail,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdatePassword(id int, newPassword string) error {
	var salt string
	err := db.Conn.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE id = $1",
		id).Scan(&salt)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrSelectSaltInUpdate,
			Time: time.Now(),
		}
	}

	_, err = db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET password = $1 WHERE id = $2",
		prof.HashPassword(newPassword, salt), id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdatePassword,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdatePhone(id int, newPhone string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET phone = $1 WHERE id = $2",
		newPhone, id)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint "+
			"\"general_user_info_phone_key\" (SQLSTATE 23505)" {
			return &errorsConst.Errors{
				Text: errorsConst.ErrUpdatePhoneRepeat,
				Time: time.Now(),
			}
		}
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdatePhone,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateAvatar(id int, newAvatar string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
		newAvatar, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdateAvatar,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateBirthday(id int, newBirthday time.Time) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE client SET date_birthday = $1 WHERE client_id = $2",
		newBirthday, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdateBirthday,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateAddress(id int, newAddress Utils.AddressCoordinates) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE address_user SET alias = $1, comment = $2, city = $3, street = $4, house = $5, floor = $6,"+
			" flat = $7, porch = $8, intercom = $9, latitude = $10, longitude = $11 WHERE client_id = $12",
		newAddress.Alias, newAddress.Comment, newAddress.City, newAddress.Street,
		newAddress.House, newAddress.Floor, newAddress.Flat, newAddress.Porch,
		newAddress.Intercom, newAddress.Coordinates.Latitude, newAddress.Coordinates.Longitude,
		id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.ErrUpdateAddress,
			Time: time.Now(),
		}
	}

	return nil
}
