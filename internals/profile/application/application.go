//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm WrapperProfileInterface
package application

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	Profile2 "2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/profile/orm"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"math"
	"strconv"
	"strings"
)

type ProfileApplicationInterface interface {
	GetProfile(id int) (*Profile2.Profile, error)
	UpdateName(id int, newName string) error
	UpdateEmail(id int, newEmail string) error
	UpdatePassword(id int, newPassword string) error
	UpdatePhone(id int, newPhone string) error
	UpdateAvatar(id int, newAvatar *Profile2.UpdateAvatar) error
	UpdateBirthday(id int, newBirthday string) error
	UpdateAddress(id int, newAddress Profile2.AddressCoordinates) error
	AddAddress(id int, newAddress Profile2.AddressCoordinates) (int, error)
	DeleteAddress(id int, addressId int) error
}

type Profile struct {
	DB ormPkg.WrapperProfileInterface
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
			Text: errPkg.PGetProfileUnknownRole,
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
			Text: errPkg.PUpdateAvatarFileNameEmpty,
		}
	}

	startExtension := strings.LastIndex(header.Filename, ".")
	if startExtension == -1 {
		return &errPkg.Errors{
			Text: errPkg.PUpdateAvatarFileWithoutExtension,
		}
	}
	extensionFile := header.Filename[startExtension:]

	fileName := strconv.Itoa(util.RandomInteger(0, math.MaxInt64))

	fileResult := "/user/" + fileName + extensionFile

	newAvatar.Avatar = fileResult
	return p.DB.UpdateAvatar(id, newAvatar, fileResult)
}

func (p *Profile) UpdateBirthday(id int, newBirthday string) error {
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
