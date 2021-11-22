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
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	role := 0

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdHostNotScan,
		}
	}
	if role != 0 {
		return "host", nil
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdClientNotScan,
		}
	}
	if role != 0 {
		return "client", nil
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdCourierNotScan,
		}
	}
	if role != 0 {
		return "courier", nil
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Alias: errPkg.PGetRoleByIdNotCommit,
		}
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*Profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var profile = Profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostHostNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileHostNotCommit,
		}
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*Profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var profile = Profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientClientNotScan,
		}
	}
	timeVoid := time.Time{}
	if timeVoid != profile.Birthday {
		err = tx.QueryRow(contextTransaction,
			"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&profile.Birthday)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.PGetProfileClientBirthdayNotScan,
			}
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileClientNotCommit,
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*Profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface) {
		tx.Rollback(contextTransaction)
	}(tx)

	var profile = Profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierCourierNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileCourierNotCommit,
		}
	}
	return &profile, nil
}

func (db *Wrapper) UpdateName(id int, newName string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateNameTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET name = $1 WHERE id = $2",
		Utils2.Sanitize(newName), id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateNameNameNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateNameNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateEmail(id int, newEmail string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateEmailTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	_, err = tx.Exec(contextTransaction,
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

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateEmailNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePassword(id int, newPassword string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var salt string
	err = tx.QueryRow(contextTransaction,
		"SELECT salt FROM general_user_info WHERE id = $1",
		id).Scan(&salt)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordSaltNotSelect,
		}
	}

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET password = $1 WHERE id = $2",
		Utils2.HashPassword(newPassword, salt), id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordPasswordNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePasswordNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePhone(id int, newPhone string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePhoneTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	if _, err := strconv.Atoi(newPhone); err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePhoneIncorrectPhoneFormat,
		}
	}

	_, err = tx.Exec(contextTransaction,
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

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdatePhoneNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAvatar(id int, newAvatar *Profile.UpdateAvatar, newFileName string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	header := newAvatar.FileHeader
	file, errTet := header.Open()
	if errTet != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarAvatarNotOpen,
		}
	}

	_, err = db.Uploader.Upload(&s3manager.UploadInput{
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

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
		newAvatar.Avatar, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarAvatarNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateBirthday(id int, newBirthday time.Time) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateBirthdayTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE client SET date_birthday = $1 WHERE client_id = $2",
		newBirthday, id)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateBirthdayBirthdayNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateBirthdayNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAddress(id int, newAddress Profile.AddressCoordinates) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAddressTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	newAddress.Sanitize()
	_, err = tx.Exec(contextTransaction,
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

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAddressNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) AddAddress(id int, newAddress Profile.AddressCoordinates) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PAddAddressNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	var idAddress int
	newAddress.Sanitize()
	err = db.Conn.QueryRow(contextTransaction,
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

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Alias: errPkg.PAddAddressNotCommit,
		}
	}

	return idAddress, nil
}

func (db *Wrapper) DeleteAddress(id int, addressId int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PAddDeleteAddressTransactionNotCreate,
		}
	}

	defer func(tx Interface.TransactionInterface, contextTransaction context.Context) {
		tx.Rollback(contextTransaction)
	}(tx, contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE address_user SET deleted = true WHERE client_id = $1 AND id = $2",
		id, addressId)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PAddDeleteAddressNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.PAddDeleteAddressNotCommit,
		}
	}

	return nil
}
