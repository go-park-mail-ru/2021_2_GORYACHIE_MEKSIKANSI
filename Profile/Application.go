package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"2021_2_GORYACHIE_MEKSIKANSI/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
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
	err := p.DB.UpdateName(id, newName)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateEmail(id int, newEmail string) error {
	err := p.DB.UpdateEmail(id, newEmail)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdatePassword(id int, newPassword string) error {
	err := p.DB.UpdatePassword(id, newPassword)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdatePhone(id int, newPhone string) error {
	err := p.DB.UpdatePhone(id, newPhone)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateAvatar(id int, newAvatar *utils.UpdateAvatar) error {
	err := p.DB.UpdateAvatar(id, newAvatar)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateBirthday(id int, newBirthday time.Time) error {
	err := p.DB.UpdateBirthday(id, newBirthday)
	if err != nil {
		return err
	}
	return nil
}

func (p *Profile) UpdateAddress(id int, newAddress utils.AddressCoordinates) error {
	err := p.DB.UpdateAddress(id, newAddress)
	if err != nil {
		return err
	}
	return nil
}


