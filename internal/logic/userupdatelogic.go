package logic

import (
	"context"
	"gitee.com/fireflylove/user-svc/internal/svc"
	"gitee.com/fireflylove/user-svc/model"
	"gitee.com/fireflylove/user-svc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *user.UserUpdateReq) (*user.Response, error) {
	u := model.User{
		Name:     in.Name,
		Avatar:   in.Avatar,
		Password: in.Password,
	}

	err := l.svcCtx.UserModel.Update(u)

	if err != nil {
		return &user.Response{Code: 1}, nil
	}

	return &user.Response{Code: 0}, nil
}
