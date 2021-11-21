package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Profile"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

type Wrapper struct {
	Conn       Interface.ConnectionInterface
	Uploader   Interface.Uploader
	NameBucket string
}

func (db *Wrapper) GetRoleById(id int) (string, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	role := 0

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdHostNotScan,
		}
	}
	if role != 0 {
		return "host", nil
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdClientNotScan,
		}
	}
	if role != 0 {
		return "client", nil
	}

	err = tx.QueryRow(context.Background(),
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdCourierNotScan,
		}
	}
	if role != 0 {
		return "courier", nil
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdNotCommit,
		}
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*Profile.Profile, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var profile = Profile.Profile{}
	err = db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostHostNotScan,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostNotCommit,
		}
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*Profile.Profile, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var profile = Profile.Profile{}
	err = tx.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientClientNotScan,
		}
	}
	timeVoid := time.Time{}
	if timeVoid != profile.Birthday {
		err = tx.QueryRow(context.Background(),
			"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&profile.Birthday)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.PGetProfileClientBirthdayNotScan,
			}
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientNotCommit,
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*Profile.Profile, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var profile = Profile.Profile{}
	err = db.Conn.QueryRow(context.Background(),
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierCourierNotScan,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierNotCommit,
		}
	}
	return &profile, nil
}

func (db *Wrapper) UpdateName(id int, newName string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET name = $1 WHERE id = $2",
		Utils2.Sanitize(newName), id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateNameNameNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) UpdateEmail(id int, newEmail string) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET email = $1 WHERE id = $2",
		Utils2.Sanitize(newEmail), id)
	if err != nil {
		textError := err.Error()
		if strings.Contains(textError, "duplicate key") {
			return &errPkg.Errors{
				Alias: errPkg.PUpdateEmailEmailRepeat,
			}
		}
		return &errPkg.Errors{
			Alias: errPkg.PUpdateEmailEmailNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePassword(id int, newPassword string) error {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var salt string
	err = db.Conn.QueryRow(context.Background(),
		"SELECT salt FROM general_user_info WHERE id = $1",
		id).Scan(&salt)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordSaltNotSelect,
		}
	}

	_, err = db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET password = $1 WHERE id = $2",
		Utils2.HashPassword(newPassword, salt), id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordPasswordNotUpdate,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePhone(id int, newPhone string) error {
	if _, err := strconv.Atoi(newPhone); err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePhoneIncorrectPhoneFormat,
		}
	}

	_, err := db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET phone = $1 WHERE id = $2",
		newPhone, id)
	if err != nil {
		errText := err.Error()
		if strings.Contains(errText, "duplicate key") {
			return &errPkg.Errors{
				Alias: errPkg.PUpdatePhonePhoneRepeat,
			}
		}
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePhonePhoneNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAvatar(id int, newAvatar *Profile.UpdateAvatar, newFileName string) error {
	header := newAvatar.FileHeader
	file, errTet := header.Open()
	if errTet != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarAvatarNotOpen,
		}
	}

	_, err := db.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(db.NameBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(newFileName),
		Body:   file,
	})
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarAvatarNotUpload,
		}
	}

	_, err = db.Conn.Exec(context.Background(),
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
		newAvatar.Avatar, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarAvatarNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) UpdateBirthday(id int, newBirthday time.Time) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE client SET date_birthday = $1 WHERE client_id = $2",
		newBirthday, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateBirthdayBirthdayNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAddress(id int, newAddress Profile.AddressCoordinates) error {
	newAddress.Sanitize()
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE address_user SET alias = $1, comment = $2, city = $3, street = $4, house = $5, floor = $6,"+
			" flat = $7, porch = $8, intercom = $9, latitude = $10, longitude = $11"+
			" WHERE client_id = $12 AND deleted = false",
		newAddress.Alias, newAddress.Comment, newAddress.City,
		newAddress.Street, newAddress.House, newAddress.Floor, newAddress.Flat,
		newAddress.Porch, newAddress.Intercom, newAddress.Coordinates.Latitude,
		newAddress.Coordinates.Longitude, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAddressAddressNotUpdate,
		}
	}

	return nil
}

func (db *Wrapper) AddAddress(id int, newAddress Profile.AddressCoordinates) (int, error) {
	var idAddress int
	newAddress.Sanitize()
	err := db.Conn.QueryRow(context.Background(),
		"INSERT INTO address_user (city, street, house, floor, flat, porch, intercom, latitude, longitude, client_id, deleted)"+
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, true) RETURNING id",
		newAddress.City, newAddress.Street, newAddress.House,
		newAddress.Floor, newAddress.Flat, newAddress.Porch,
		newAddress.Intercom, newAddress.Coordinates.Latitude,
		newAddress.Coordinates.Longitude, id).Scan(&idAddress)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PAddAddressAddressNotAdd,
		}
	}

	return idAddress, nil
}

func (db *Wrapper) DeleteAddress(id int, addressId int) error {
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE address_user SET deleted = true WHERE client_id = $1 AND id = $2",
		id, addressId)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PAddDeleteAddressNotDelete,
		}
	}

	return nil
}
