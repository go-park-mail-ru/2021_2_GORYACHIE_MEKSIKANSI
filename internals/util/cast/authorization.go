package cast

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func CastDefenseToDefenseProto(d *Utils2.Defense) *authProto.Defense {
	var p authProto.Defense
	p.DateLife = timestamp.New(d.DateLife)
	p.XCsrfToken = d.CsrfToken
	p.SessionId = d.SessionId
	return &p
}

func CastDefenseResponseProtoToDefense(p *authProto.DefenseResponse) *Utils2.Defense {
	var d Utils2.Defense
	d.DateLife = p.Defense.DateLife.AsTime()
	d.CsrfToken = p.Defense.XCsrfToken
	d.SessionId = p.Defense.SessionId
	return &d
}

func CastAuthorizationToAuthorizationProto(a *auth.Authorization) *authProto.Authorization {
	var p authProto.Authorization
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	return &p
}

func CastRegistrationRequestToRegistrationRequestProto(a *auth.RegistrationRequest) *authProto.RegistrationRequest {
	var p authProto.RegistrationRequest
	p.Phone = a.Phone
	p.Email = a.Email
	p.Password = a.Password
	p.TypeUser = a.TypeUser
	p.Name = a.Name
	return &p
}
