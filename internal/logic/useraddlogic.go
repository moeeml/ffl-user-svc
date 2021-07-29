package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user-svc/model"
	"user-svc/tool"

	"user-svc/internal/svc"
	"user-svc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *user.UserAddReq) (*user.UserAddRsp, error) {

	m := l.svcCtx.UserModel
	checkAccount, _ := m.FindOneByAccount(in.Account)
	if checkAccount != nil {
		return &user.UserAddRsp{
			Code:    10001,
			Message: tool.ErrorCode[10001],
		}, nil
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	r, err := m.Insert(model.User{
		Account:    in.Account,
		Password:   string(pwd),
		Name:       in.Name,
		Avatar:     in.Avatar,
		Status:     tool.UserStatusNormal,
		CreateTime: time.Now().Unix(),
	})

	if err != nil {
		return &user.UserAddRsp{Code: 1}, nil
	}

	uid, _ := r.LastInsertId()

	return &user.UserAddRsp{
		Code: 0,
		Uid:  uid,
	}, nil
}
