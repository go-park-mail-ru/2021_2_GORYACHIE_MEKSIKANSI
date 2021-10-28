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
	ProfileUser interface{} `json:"user"`
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

type UpdateAddress struct {
	Address AddressCoordinates `json:"address"`
}

type AddressCoordinates struct {
	Coordinates Coordinates `json:"coordinates"`
	Alias       string      `json:"alias"`
	City        string      `json:"city"`
	Street      string      `json:"street,omitempty"`
	House       string      `json:"house,omitempty"`
	Flat        int         `json:"flat,omitempty"`
	Porch       int         `json:"porch,omitempty"`
	Floor       int         `json:"floor,omitempty"`
	Intercom    string      `json:"intercom,omitempty"`
	Comment     string      `json:"comment,omitempty"`
}

type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
