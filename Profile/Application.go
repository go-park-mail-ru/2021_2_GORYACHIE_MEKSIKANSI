package Profile
import (
	profile "2021_2_GORYACHIE_MEKSIKANSI/Utils"
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
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
