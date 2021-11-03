package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	profile "2021_2_GORYACHIE_MEKSIKANSI/Utils"

	"time"
)

func GetProfile(db profile.WrapperProfile, id int) (*profile.Profile, error) {
	role, err := db.GetRoleById(id)
	if err != nil {
		return nil, err
	}

	var result *profile.Profile
	switch role {
	case "client":
		result, err = db.GetProfileClient(id)
	case "courier":
		result, err = db.GetProfileCourier(id)
	case "host":
		result, err = db.GetProfileHost(id)
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

func UpdateName(db profile.WrapperProfile, id int, newName string) error {
	err := db.UpdateName(id, newName)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmail(db profile.WrapperProfile, id int, newEmail string) error {
	err := db.UpdateEmail(id, newEmail)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePassword(db profile.WrapperProfile, id int, newPassword string) error {
	err := db.UpdatePassword(id, newPassword)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePhone(db profile.WrapperProfile, id int, newPhone string) error {
	err := db.UpdatePhone(id, newPhone)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAvatar(db profile.WrapperProfile, id int, newAvatar string) error {

	err := db.UpdateAvatar(id, newAvatar)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBirthday(db profile.WrapperProfile, id int, newBirthday time.Time) error {
	err := db.UpdateBirthday(id, newBirthday)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAddress(db profile.WrapperProfile, id int, newAddress profile.AddressCoordinates) error {
	err := db.UpdateAddress(id, newAddress)
	if err != nil {
		return err
	}
	return nil
}
