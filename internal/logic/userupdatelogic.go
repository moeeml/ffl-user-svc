package logic

import (
	"context"
	"gitee.com/fireflylove/user-svc/internal/svc"
	"gitee.com/fireflylove/user-svc/model"
	"gitee.com/fireflylove/user-svc/user"
	"golang.org/x/crypto/bcrypt"

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

	var u model.User
	var db = l.svcCtx.DB
	db.First(&u, in.Id)

	u.Name = in.Name
	u.Avatar = in.Avatar

	if len(in.Password) > 0 {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		u.Password = string(pwd)
	}

	if r := db.Save(u); r.Error != nil {
		return &user.Response{Code: 1}, nil
	}

	return &user.Response{Code: 0}, nil
}
