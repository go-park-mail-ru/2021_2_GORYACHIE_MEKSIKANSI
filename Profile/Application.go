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

//func UpdateName(db Wrapper, cookie Defense, id int, name string) error {
//	err := db.updateName(id, name)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func UpdateEmail(db Wrapper, cookie Defense, id int, email string) error {
//	err := db.updateEmail(id, email)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func UpdatePassword(db Wrapper, cookie Defense, id int, password string) error {
//	err := db.updatePassword(id, password)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func UpdateAdditionalInfo(db Wrapper, cookie Defense, id int, phone string) error {
//	err := db.updateAdditionalInfo(id, phone)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func UpdateAvatar(db Wrapper, cookie Defense, id int, avatar string) error {
//	err := db.updateAvatar(id, avatar)
//	if err != nil {
//		return err
//	}
//	return nil
//}
