package service

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/authorization/proto"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func CastDefenseToDefenseProto(d *authPkg.Defense) *authProto.Defense {
	var p authProto.Defense
	p.DateLife = timestamp.New(d.DateLife)
	p.XCsrfToken = d.CsrfToken
	p.SessionId = d.SessionId
	return &p
}

func CastDefenseProtoToDefense(p *authProto.Defense) *authPkg.Defense {
	var d authPkg.Defense
	d.DateLife = p.DateLife.AsTime()
	d.CsrfToken = p.XCsrfToken
	d.SessionId = p.SessionId
	return &d
}

func CastAuthorizationProtoToAuthorization(a *authProto.Authorization) *authPkg.Authorization {
	var p authPkg.Authorization
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	return &p
}

func CastRegistrationRequestProtoToRegistrationRequest(p *authProto.RegistrationRequest) *authPkg.RegistrationRequest {
	var a authPkg.RegistrationRequest
	a.Phone = p.Phone
	a.Email = p.Email
	a.Password = p.Password
	a.TypeUser = p.TypeUser
	a.Name = p.Name
	return &a
}
