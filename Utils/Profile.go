package Utils

import "time"

type Profile struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday,omitempty"`
}

type ProfileResponse struct {
	ProfileUser interface{} `json:"profile"`
}

type UpdateName struct {
	Name string `json:"name"`
}

type UpdateEmail struct {
	Email string `json:"email"`
}

type UpdatePassword struct {
	Password string `json:"password"`
}

type UpdatePhone struct {
	Phone string `json:"Phone"`
}

type UpdateAvatar struct {
	Avatar string `json:"avatar"` // TODO(N): проверить тип
}

type UpdateBirthday struct {
	Birthday time.Time `json:"birthday"`
}
