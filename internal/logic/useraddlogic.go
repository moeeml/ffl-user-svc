package logic

import (
	"context"
	"errors"
	"gitee.com/fireflylove/user-svc/common"
	"gitee.com/fireflylove/user-svc/internal/svc"
	"gitee.com/fireflylove/user-svc/model"
	"gitee.com/fireflylove/user-svc/user"
	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	var u model.User
	tx := l.svcCtx.DB.Begin()

	r := tx.Where(&model.User{Account: in.Account}).First(&u)
	if !errors.Is(r.Error, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return &user.UserAddRsp{
			Code:    10002,
			Message: common.ErrorCode[10002],
		}, nil
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	o := &model.User{
		Account:  in.Account,
		Password: string(pwd),
		Name:     in.Name,
		Avatar:   in.Avatar,
	}

	ir := tx.Create(o)
	if ir.Error != nil {
		tx.Rollback()
		return &user.UserAddRsp{Code: 1}, nil
	}

	tx.Commit()

	return &user.UserAddRsp{
		Code: 0,
		Uid:  uint64(o.ID),
	}, nil
}
