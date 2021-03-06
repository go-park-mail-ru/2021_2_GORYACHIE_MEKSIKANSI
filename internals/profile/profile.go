package profile

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"mime/multipart"
)

const (
	PhoneLen = 11
)

type Profile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Birthday string `json:"birthday,omitempty"`
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
	Avatar     string
	FileHeader *multipart.FileHeader
}

type UpdateAvatarRequest struct {
	PathImg string `json:"img"`
}

type UpdateBirthday struct {
	Birthday string `json:"birthday"`
}

type UpdateAddress struct {
	Address AddressCoordinates `json:"address"`
}

type AddressCoordinates struct {
	Coordinates Coordinates `json:"coordinates"`
	Alias       string      `json:"alias,omitempty"`
	City        string      `json:"city"`
	Street      string      `json:"street,omitempty"`
	House       string      `json:"house,omitempty"`
	Flat        string      `json:"flat,omitempty"`
	Porch       int         `json:"porch,omitempty"`
	Floor       int         `json:"floor,omitempty"`
	Intercom    string      `json:"intercom,omitempty"`
	Comment     string      `json:"comment,omitempty"`
}

type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (a *AddressCoordinates) Sanitize() {
	a.Alias = util.Sanitize(a.Alias)
	a.Comment = util.Sanitize(a.Comment)
	a.City = util.Sanitize(a.City)
	a.Street = util.Sanitize(a.Street)
	a.House = util.Sanitize(a.House)
}

func Sanitize(str string) string {
	return util.Sanitize(str)
}
