package Profile

func GetProfile(db Wrapper, id int) (*Profile, error) {
	role, err := db.getRoleById(id)
	if err != nil {
		return nil, err
	}

	var result *Profile
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
