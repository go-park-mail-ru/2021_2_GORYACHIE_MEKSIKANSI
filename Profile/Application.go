package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"

	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"math"
	"strconv"
	"strings"
	"time"
)

type Profile struct {
	DB Interfaces.WrapperProfile
}

func (p *Profile) GetProfile(id int) (*utils.Profile, error) {
	role, err := p.DB.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	var result *utils.Profile
	switch role {
	case "client":
		result, err = p.DB.GetProfileClient(id)
	case "courier":
		result, err = p.DB.GetProfileCourier(id)
	case "host":
		result, err = p.DB.GetProfileHost(id)
	default:
		return nil, &errorsConst.Errors{
			Text: errorsConst.PGetProfileUnknownRole,
			Time: time.Now(),
		}
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Profile) UpdateName(id int, newName string) error {
	return p.DB.UpdateName(id, newName)
}

func (p *Profile) UpdateEmail(id int, newEmail string) error {
	return p.DB.UpdateEmail(id, newEmail)
}

func (p *Profile) UpdatePassword(id int, newPassword string) error {
	return p.DB.UpdatePassword(id, newPassword)
}

func (p *Profile) UpdatePhone(id int, newPhone string) error {
	return p.DB.UpdatePhone(id, newPhone)
}

func (p *Profile) UpdateAvatar(id int, newAvatar *utils.UpdateAvatar) error {
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

	newAvatar.Avatar = fileResult
	return p.DB.UpdateAvatar(id, newAvatar, fileResult)
}

func (p *Profile) UpdateBirthday(id int, newBirthday time.Time) error {
	return p.DB.UpdateBirthday(id, newBirthday)
}

func (p *Profile) UpdateAddress(id int, newAddress utils.AddressCoordinates) error {
	return p.DB.UpdateAddress(id, newAddress)
}

func (p *Profile) AddAddress(id int, newAddress utils.AddressCoordinates) (int, error) {
	return p.DB.AddAddress(id, newAddress)
}

func (p *Profile) DeleteAddress(id int, addressId int) error {
	return p.DB.DeleteAddress(id, addressId)
}
