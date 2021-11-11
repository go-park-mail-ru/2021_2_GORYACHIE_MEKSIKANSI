package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

type Wrapper struct {
	Conn       Interfaces.ConnectionInterface
	Uploader   Interfaces.Uploader
	NameBucket string
}

func (db *Wrapper) GetRoleById(id int) (string, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return "", &errorsConst.Errors{
			Text: errorsConst.PGetRoleByIdTransactionNotCreate,
			Time: time.Now(),
		}
	}

	role := 0

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && strings.Contains(err.Error(), "no rows") != true {
		return "", &errorsConst.Errors{
			Text: errorsConst.PGetRoleByIdHostNotScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "host", nil
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && strings.Contains(err.Error(), "no rows") != true {
		return "", &errorsConst.Errors{
			Text: errorsConst.PGetRoleByIdClientNotScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "client", nil
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && strings.Contains(err.Error(), "no rows") != true {
		return "", &errorsConst.Errors{
			Text: errorsConst.PGetRoleByIdCourierNotScan,
			Time: time.Now(),
		}
	}
	if role != 0 {
		return "courier", nil
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return "", &errorsConst.Errors{
			Text: errorsConst.PGetRoleByIdNotCommit,
			Time: time.Now(),
		}
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*utils.Profile, error) {
	var profile = utils.Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileHostHostNotScan,
			Time: time.Now(),
		}
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*utils.Profile, error) {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileClientTransactionNotCreate,
			Time: time.Now(),
		}
	}

	var profile = utils.Profile{}
	err = db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileClientClientNotScan,
			Time: time.Now(),
		}
	}
	timeVoid := time.Time{}
	if timeVoid != profile.Birthday {
		err = db.Conn.QueryRow(context.Background(),
			"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&profile.Birthday)
		if err != nil {
			return nil, &errorsConst.Errors{
				Text: errorsConst.PGetProfileClientBirthdayNotScan,
				Time: time.Now(),
			}
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileClientNotCommit,
			Time: time.Now(),
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*utils.Profile, error) {
	var profile = utils.Profile{}
	err := db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileCourierCourierNotScan,
			Time: time.Now(),
		}
	}
	return &profile, nil
}

func (db *Wrapper) UpdateName(id int, newName string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET name = $1 WHERE id = $2",
		Utils.Sanitize(newName), id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateNameNameNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateEmail(id int, newEmail string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET email = $1 WHERE id = $2",
		Utils.Sanitize(newEmail), id)
	if err != nil {
		textError := err.Error()
		println(textError)
		if textError == "ERROR: duplicate key value violates unique constraint "+
			"\"general_user_info_email_key\" (SQLSTATE 23505)" {
			return &errorsConst.Errors{
				Text: errorsConst.PUpdateEmailEmailRepeat,
				Time: time.Now(),
			}
		}
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateEmailEmailNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdatePassword(id int, newPassword string) error {
	tx, err := db.Conn.Begin(context.Background())

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePasswordTransactionNotCreate,
			Time: time.Now(),
		}
	}

	var salt string
	err = db.Conn.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE id = $1",
		id).Scan(&salt)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePasswordSaltNotSelect,
			Time: time.Now(),
		}
	}

	_, err = db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET password = $1 WHERE id = $2",
		utils.HashPassword(newPassword, salt), id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePasswordPasswordNotUpdate,
			Time: time.Now(),
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePasswordNotCommit,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdatePhone(id int, newPhone string) error {
	if _, err := strconv.Atoi(newPhone); err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePhoneIncorrectPhoneFormat,
			Time: time.Now(),
		}
	}

	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET phone = $1 WHERE id = $2",
		newPhone, id)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint "+
			"\"general_user_info_phone_key\" (SQLSTATE 23505)" {
			return &errorsConst.Errors{
				Text: errorsConst.PUpdatePhonePhoneRepeat,
				Time: time.Now(),
			}
		}
		return &errorsConst.Errors{
			Text: errorsConst.PUpdatePhonePhoneNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateAvatar(id int, newAvatar *Utils.UpdateAvatar, newFileName string) error {
	header := newAvatar.FileHeader
	file, errTet := header.Open()
	if errTet != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarAvatarNotOpen,
			Time: time.Now(),
		}
	}

	_, err := db.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(db.NameBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(newFileName),
		Body:   file,
	})
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarAvatarNotUpload,
			Time: time.Now(),
		}
	}

	_, err = db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
		newAvatar.Avatar, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarAvatarNotUpdate,
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
			Text: errorsConst.PUpdateBirthdayBirthdayNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}

func (db *Wrapper) UpdateAddress(id int, newAddress Utils.AddressCoordinates) error {
	newAddress.Sanitize()
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE address_user SET alias = $1, comment = $2, city = $3, street = $4, house = $5, floor = $6,"+
			" flat = $7, porch = $8, intercom = $9, latitude = $10, longitude = $11 WHERE client_id = $12",
		newAddress.Alias, newAddress.Comment, newAddress.City,
		newAddress.Street, newAddress.House, newAddress.Floor, newAddress.Flat,
		newAddress.Porch, newAddress.Intercom, newAddress.Coordinates.Latitude,
		newAddress.Coordinates.Longitude, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAddressAddressNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}
