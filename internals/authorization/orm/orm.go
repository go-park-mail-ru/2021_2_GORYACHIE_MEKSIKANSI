//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm WrapperAuthorizationInterface,ConnectAuthServiceInterface
package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util/cast"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"io/ioutil"
)

const (
	APP_SECRET = "8e0564128e0564128e689879548e7ff1ff88e058e056412efc6c1f72a21a320b854ad06"
	API_URL    = "https://api.vk.com/method/users.get?fields=contacts&access_token=%s&v=5.131"
)

type WrapperAuthorizationInterface interface {
	SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error)
	Login(login *authorization.Authorization) (*Utils2.Defense, error)
	Logout(CSRF string) (string, error)
	NewCSRFWebsocket(id int) (string, error)
	AuthVK(code string) (*Utils2.Defense, error)
}

type ConnectAuthServiceInterface interface {
	SignUp(ctx context.Context, in *authProto.RegistrationRequest, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Login(ctx context.Context, in *authProto.Authorization, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
	Logout(ctx context.Context, in *authProto.CSRF, opts ...grpc.CallOption) (*authProto.CSRFResponse, error)
	NewCSRFWebsocket(ctx context.Context, client *authProto.IdClient, opts ...grpc.CallOption) (*authProto.WebsocketResponse, error)
	AuthVK(ctx context.Context, client *authProto.PartSignUp, opts ...grpc.CallOption) (*authProto.DefenseResponse, error)
}

type Wrapper struct {
	Conn   ConnectAuthServiceInterface
	Ctx    context.Context
	VKConn oauth2.Config
}

func (w *Wrapper) SignUp(signup *authorization.RegistrationRequest) (*Utils2.Defense, error) {
	result, err := w.Conn.SignUp(w.Ctx, cast.CastRegistrationRequestToRegistrationRequestProto(signup))
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, &errPkg.Errors{Text: result.Error}
	}
	return cast.CastDefenseResponseProtoToDefense(result), nil
}

func (w *Wrapper) Login(login *authorization.Authorization) (*Utils2.Defense, error) {
	response, err := w.Conn.Login(w.Ctx, cast.CastAuthorizationToAuthorizationProto(login))
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, &errPkg.Errors{Text: response.Error}
	}
	return cast.CastDefenseResponseProtoToDefense(response), nil
}

func (w *Wrapper) Logout(CSRF string) (string, error) {
	var csrfToken authProto.CSRF
	csrfToken.XCsrfToken = CSRF
	logout, err := w.Conn.Logout(w.Ctx, &csrfToken)
	if err != nil {
		return "", err
	}
	if logout.Error != "" {
		return "", &errPkg.Errors{Text: logout.Error}
	}
	return logout.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) NewCSRFWebsocket(id int) (string, error) {
	var idClient authProto.IdClient
	idClient.ClientId = int64(id)
	websocket, err := w.Conn.NewCSRFWebsocket(w.Ctx, &idClient)
	if err != nil {
		return "", err
	}
	if websocket.Error != "" {
		return "", &errPkg.Errors{
			Text: websocket.Error,
		}
	}
	return websocket.Websocket, nil
}

type Response struct {
	Response []struct {
		FirstName string `json:"first_name"`
	}
}

func (w *Wrapper) AuthVK(code string) (*Utils2.Defense, error) {
	token, err := w.VKConn.Exchange(w.Ctx, code)
	if err != nil {
		return nil, nil
	}
	email := token.Extra("email").(string)

	client := w.VKConn.Client(w.Ctx, token)

	resp, err := client.Get(fmt.Sprintf(API_URL, APP_SECRET))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	data := &Response{}
	json.Unmarshal(body, data)
	signUp := &authProto.PartSignUp{
		Email: email,
		Name:  data.Response[0].FirstName,
	}
	result, err := w.Conn.AuthVK(w.Ctx, signUp)
	return cast.CastDefenseResponseProtoToDefense(result), nil
}
