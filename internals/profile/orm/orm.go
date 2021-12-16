//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm WrapperProfileInterface,ConnectionInterface,UploaderInterface,TransactionInterface
package orm

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"strconv"
	"strings"
	"time"
)

type WrapperProfileInterface interface {
	GetRoleById(id int) (string, error)
	GetProfileClient(id int) (*profile.Profile, error)
	GetProfileHost(id int) (*profile.Profile, error)
	GetProfileCourier(id int) (*profile.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *profile.UpdateAvatar, newFileName string) error
	UpdateBirthday(id int, newBirthday string) error
	UpdateAddress(id int, newAddress profile.AddressCoordinates) error
	AddAddress(id int, newAddress profile.AddressCoordinates) (int, error)
	DeleteAddress(id int, addressId int) error
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type UploaderInterface interface {
	Upload(input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

type Wrapper struct {
	Conn       ConnectionInterface
	Uploader   UploaderInterface
	NameBucket string
}

func (db *Wrapper) GetRoleById(id int) (string, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.PGetRoleByIdTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	role := 0

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM host WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Text: errPkg.PGetRoleByIdHostNotScan,
		}
	}
	if role != 0 {
		err = tx.Commit(contextTransaction)
		if err != nil {
			return "", &errPkg.Errors{
				Text: errPkg.PGetRoleByIdNotCommit,
			}
		}
		return "host", nil
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM client WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Text: errPkg.PGetRoleByIdClientNotScan,
		}
	}
	if role != 0 {
		err = tx.Commit(contextTransaction)
		if err != nil {
			return "", &errPkg.Errors{
				Text: errPkg.PGetRoleByIdNotCommit,
			}
		}
		return "client", nil
	}

	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM courier WHERE client_id = $1", id).Scan(&role)
	if err != nil && err != pgx.ErrNoRows {
		return "", &errPkg.Errors{
			Text: errPkg.PGetRoleByIdCourierNotScan,
		}
	}
	if role != 0 {
		err = tx.Commit(contextTransaction)
		if err != nil {
			return "", &errPkg.Errors{
				Text: errPkg.PGetRoleByIdNotCommit,
			}
		}
		return "courier", nil
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return "", &errPkg.Errors{
			Text: errPkg.PGetRoleByIdNotCommit,
		}
	}

	return "", nil
}

func (db *Wrapper) GetProfileHost(id int) (*profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileHostTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var profile = profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileHostHostNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileHostNotCommit,
		}
	}

	return &profile, err
}

func (db *Wrapper) GetProfileClient(id int) (*profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileClientTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var profile = profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileClientClientNotScan,
		}
	}

	var birthday *time.Time
	err = tx.QueryRow(contextTransaction,
		"SELECT date_birthday FROM client WHERE client_id = $1", id).Scan(&birthday)
	if birthday != nil {
		profile.Birthday, _ = Utils2.FormatDate(*birthday)
	}
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileClientBirthdayNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileClientNotCommit,
		}
	}

	return &profile, nil
}

func (db *Wrapper) GetProfileCourier(id int) (*profile.Profile, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileCourierTransactionNotCreate,
		}
	}

	defer func(tx TransactionInterface) {
		tx.Rollback(contextTransaction)
	}(tx)

	var profile = profile.Profile{}
	err = tx.QueryRow(contextTransaction,
		"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", id).Scan(
		&profile.Email, &profile.Name, &profile.Avatar, &profile.Phone)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileCourierCourierNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.PGetProfileCourierNotCommit,
		}
	}
	return &profile, nil
}

func (db *Wrapper) UpdateName(id int, newName string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateNameTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET name = $1 WHERE id = $2",
		Utils2.Sanitize(newName), id)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateNameNameNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateNameNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateEmail(id int, newEmail string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateEmailTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET email = $1 WHERE id = $2",
		Utils2.Sanitize(newEmail), id)
	if err != nil {
		textError := err.Error()
		if strings.Contains(textError, "duplicate key") {
			return &errPkg.Errors{
				Text: errPkg.PUpdateEmailEmailRepeat,
			}
		}
		return &errPkg.Errors{
			Text: errPkg.PUpdateEmailEmailNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateEmailNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePassword(id int, newPassword string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePasswordTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var salt string
	err = tx.QueryRow(contextTransaction,
		"SELECT salt FROM general_user_info WHERE id = $1",
		id).Scan(&salt)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePasswordSaltNotSelect,
		}
	}

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET password = $1 WHERE id = $2",
		Utils2.HashPassword(newPassword, salt), id)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePasswordPasswordNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePasswordNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdatePhone(id int, newPhone string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePhoneTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	profile.Sanitize(newPhone)

	if newPhone != "" && newPhone[0] == '+' && newPhone[1] == '7' {
		s := []rune(newPhone)
		s[1] = '8'
		newPhone = string(s[1:])
	}

	if _, err := strconv.Atoi(newPhone); err != nil || len(newPhone) != profile.PhoneLen {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePhoneIncorrectPhoneFormat,
		}
	}

	s := []rune(newPhone)
	s[0] = '8'
	newPhone = string(s)

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET phone = $1 WHERE id = $2",
		newPhone, id)
	if err != nil {
		errText := err.Error()
		if strings.Contains(errText, "duplicate key") {
			return &errPkg.Errors{
				Text: errPkg.PUpdatePhonePhoneRepeat,
			}
		}
		return &errPkg.Errors{
			Text: errPkg.PUpdatePhonePhoneNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdatePhoneNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAvatar(id int, newAvatar *profile.UpdateAvatar, newFileName string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAvatarTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	header := newAvatar.FileHeader
	file, errTet := header.Open()
	if errTet != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAvatarAvatarNotOpen,
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
			Text: errPkg.PUpdateAvatarAvatarNotUpload,
		}
	}

	_, err = tx.Exec(contextTransaction,
		"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
		newAvatar.Avatar, id)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAvatarAvatarNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAvatarNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateBirthday(id int, newBirthday string) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateBirthdayTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)
	layout := "02.01.2006"
	birthday, err := time.Parse(layout, newBirthday)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateBirthdayNotParse,
		}
	}
	_, err = tx.Exec(contextTransaction,
		"UPDATE client SET date_birthday = $1 WHERE client_id = $2",
		birthday, id)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateBirthdayBirthdayNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateBirthdayNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) UpdateAddress(id int, newAddress profile.AddressCoordinates) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAddressTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

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
			Text: errPkg.PUpdateAddressAddressNotUpdate,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAddressNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) AddAddress(id int, newAddress profile.AddressCoordinates) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.PAddAddressNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

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
			Text: errPkg.PAddAddressAddressNotAdd,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.PAddAddressNotCommit,
		}
	}

	return idAddress, nil
}

func (db *Wrapper) DeleteAddress(id int, addressId int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PAddDeleteAddressTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"UPDATE address_user SET deleted = true WHERE client_id = $1 AND id = $2",
		id, addressId)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PAddDeleteAddressNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.PAddDeleteAddressNotCommit,
		}
	}

	return nil
}
