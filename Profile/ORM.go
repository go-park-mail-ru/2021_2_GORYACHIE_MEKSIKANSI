package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"math"
	"strconv"
	"strings"
	"time"
)

type Wrapper struct {
	Conn Interfaces.ConnectionInterface
	Uploader Interfaces.Uploader
	NameBucket string
}

func (db *Wrapper) GetRoleById(id int) (string, error) {
	role := 0

	err := db.Conn.QueryRow(context.Background(),
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

	err = db.Conn.QueryRow(context.Background(),
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

	err = db.Conn.QueryRow(context.Background(),
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
	var profile = utils.Profile{}
	err := db.Conn.QueryRow(context.Background(),
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
	var salt string
	err := db.Conn.QueryRow(context.Background(),
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

func (db *Wrapper) UpdateAvatar(id int, newAvatar *Utils.UpdateAvatar) error {
	header := newAvatar.FileHeader
	if header.Filename == "" {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarFileNameEmpty,
			Time: time.Now(),
		}
	}
	startExtension := strings.LastIndex(header.Filename, ".")
	if startExtension == -1 {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarFileWithoutExtension,
			Time: time.Now(),
		}
	}
	extensionFile := header.Filename[startExtension:]

	fileName := strconv.Itoa(utils.RandomInteger(0, math.MaxInt64))

	fileResult := "/user/" + fileName + extensionFile
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
		Key:    aws.String(fileResult),
		Body:   file,
	})
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAvatarAvatarNotUpload,
			Time: time.Now(),
		}
	}

	newAvatar.Avatar = fileResult

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
	_, err := db.Conn.Exec(context.Background(),
		"UPDATE address_user SET alias = $1, comment = $2, city = $3, street = $4, house = $5, floor = $6,"+
			" flat = $7, porch = $8, intercom = $9, latitude = $10, longitude = $11 WHERE client_id = $12",
		Utils.Sanitize(newAddress.Alias), Utils.Sanitize(newAddress.Comment), Utils.Sanitize(newAddress.City),
		Utils.Sanitize(newAddress.Street), Utils.Sanitize(newAddress.House), newAddress.Floor, newAddress.Flat,
		newAddress.Porch, Utils.Sanitize(newAddress.Intercom), newAddress.Coordinates.Latitude,
		newAddress.Coordinates.Longitude, id)
	if err != nil {
		return &errorsConst.Errors{
			Text: errorsConst.PUpdateAddressAddressNotUpdate,
			Time: time.Now(),
		}
	}

	return nil
}
