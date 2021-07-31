// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"gitee.com/fireflylove/user-svc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	UserDetail    = user.UserDetail
	UserInfoRsp   = user.UserInfoRsp
	UserListReq   = user.UserListReq
	Response      = user.Response
	UserAddReq    = user.UserAddReq
	UserAddRsp    = user.UserAddRsp
	UserUpdateReq = user.UserUpdateReq
	UserInfo      = user.UserInfo
	UserInfoReq   = user.UserInfoReq
	UserListRsp   = user.UserListRsp

	User interface {
		UserAdd(ctx context.Context, in *UserAddReq) (*UserAddRsp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq) (*Response, error)
		UserInfo(ctx context.Context, in *UserInfoReq) (*UserInfoRsp, error)
		UserList(ctx context.Context, in *UserListReq) (*UserListRsp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserAdd(ctx context.Context, in *UserAddReq) (*UserAddRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserAdd(ctx, in)
}

func (m *defaultUser) UserUpdate(ctx context.Context, in *UserUpdateReq) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserUpdate(ctx, in)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoReq) (*UserInfoRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in)
}

func (m *defaultUser) UserList(ctx context.Context, in *UserListReq) (*UserListRsp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserList(ctx, in)
}
