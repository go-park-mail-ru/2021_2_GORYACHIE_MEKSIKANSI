package Profile

func GetProfile(db Wrapper, id int) (Profile, error) {
	role, err := db.getRoleById(id)
	if err != nil {
		return Profile{}, err
	}

	var result Profile
	switch role {
	case "client":
		result, err = db.GetProfileClient(id)
	case "courier":
		result, err = db.GetProfileCourier(id)
	case "host":
		result, err = db.GetProfileHost(id)
	default:
		if err != nil {
			return result, err
		}
	}
	if err != nil {
		return Profile{}, err
	}

	return result, nil
}
