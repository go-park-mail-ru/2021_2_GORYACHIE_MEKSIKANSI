//go:generate easyjson -no_std_marshalers authorization.go
package authorization

//easyjson:json
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

//easyjson:json
type Authorization struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

//easyjson:json
type Result struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}

type WebSocket struct {
	Socket interface{} `json:"web_socket"`
}

type KeyWebSocket struct {
	Key string `json:"key"`
}

type WebSocketAction struct {
	Action string         `json:"action"`
	Order  WebSocketOrder `json:"order"`
}

type WebSocketOrder struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

func UserConvertRegistration(signUpAll *RegistrationRequest) *User {
	user := User{
		TypeUser: signUpAll.TypeUser,
		Name:     signUpAll.Name,
		Email:    signUpAll.Email,
		Phone:    signUpAll.Phone,
	}
	return &user
}
