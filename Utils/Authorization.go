package Utils

type RegistrationRequest struct {
	TypeUser string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegistrationResponse struct {
	User interface{} `json:"user"`
}

type User struct {
	TypeUser string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Authorization struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Result struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}
