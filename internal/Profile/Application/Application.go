package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	Profile2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"math"
	"strconv"
	"strings"
	"time"
)

type Profile struct {
	DB Interface.WrapperProfile
}

func (p *Profile) GetProfile(id int) (*Profile2.Profile, error) {
	role, err := p.DB.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	var result *Profile2.Profile
	switch role {
	case "client":
		result, err = p.DB.GetProfileClient(id)
	case "courier":
		result, err = p.DB.GetProfileCourier(id)
	case "host":
		result, err = p.DB.GetProfileHost(id)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.PGetProfileUnknownRole,
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

func (p *Profile) UpdateAvatar(id int, newAvatar *Profile2.UpdateAvatar) error {
	header := newAvatar.FileHeader
	if header.Filename == "" {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarFileNameEmpty,
		}
	}

	startExtension := strings.LastIndex(header.Filename, ".")
	if startExtension == -1 {
		return &errPkg.Errors{
			Alias: errPkg.PUpdateAvatarFileWithoutExtension,
		}
	}
	extensionFile := header.Filename[startExtension:]

	fileName := strconv.Itoa(Util.RandomInteger(0, math.MaxInt64))

	fileResult := "/user/" + fileName + extensionFile

	newAvatar.Avatar = fileResult
	return p.DB.UpdateAvatar(id, newAvatar, fileResult)
}

func (p *Profile) UpdateBirthday(id int, newBirthday time.Time) error {
	return p.DB.UpdateBirthday(id, newBirthday)
}

func (p *Profile) UpdateAddress(id int, newAddress Profile2.AddressCoordinates) error {
	return p.DB.UpdateAddress(id, newAddress)
}

func (p *Profile) AddAddress(id int, newAddress Profile2.AddressCoordinates) (int, error) {
	return p.DB.AddAddress(id, newAddress)
}

func (p *Profile) DeleteAddress(id int, addressId int) error {
	return p.DB.DeleteAddress(id, addressId)
}
