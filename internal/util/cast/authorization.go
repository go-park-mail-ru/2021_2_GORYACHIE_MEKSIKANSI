package cast

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/authorization/proto"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func CastDefenseToDefenseProto(d *Utils2.Defense) *authProto.Defense {
	var p authProto.Defense
	p.DateLife = timestamp.New(d.DateLife)
	p.XCsrfToken = d.CsrfToken
	p.SessionId = d.SessionId
	return &p
}

func CastDefenseProtoToDefense(p *authProto.Defense) *Utils2.Defense {
	var d Utils2.Defense
	d.DateLife = p.DateLife.AsTime()
	d.CsrfToken = p.XCsrfToken
	d.SessionId = p.SessionId
	return &d
}

func CastDefenseResponseProtoToDefense(p *authProto.DefenseResponse) *Utils2.Defense {
	var d Utils2.Defense
	d.DateLife = p.Defense.DateLife.AsTime()
	d.CsrfToken = p.Defense.XCsrfToken
	d.SessionId = p.Defense.SessionId
	return &d
}

func CastAuthorizationToAuthorizationProto(a *authorization.Authorization) *authProto.Authorization {
	var p authProto.Authorization
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	return &p
}

func CastAuthorizationProtoToAuthorization(a *authProto.Authorization) *authorization.Authorization {
	var p authorization.Authorization
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	return &p
}

func CastRegistrationRequestToRegistrationRequestProto(a *authorization.RegistrationRequest) *authProto.RegistrationRequest {
	var p authProto.RegistrationRequest
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	p.TypeUser = a.TypeUser
	p.Name = a.Name
	return &p
}

func CastRegistrationRequestProtoToRegistrationRequest(p *authProto.RegistrationRequest) *authorization.RegistrationRequest {
	var a authorization.RegistrationRequest
	a.Phone = p.Phone
	a.Email = p.Email
	a.Password = p.Password
	a.TypeUser = p.TypeUser
	a.Name = p.Name
	return &a
}
