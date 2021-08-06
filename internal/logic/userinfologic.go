package logic

import (
	"context"
	"gitee.com/fireflylove/user-svc/model"
	"gitee.com/fireflylove/user-svc/tool"
	"gorm.io/gorm"

	"gitee.com/fireflylove/user-svc/internal/svc"
	"gitee.com/fireflylove/user-svc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoRsp, error) {
	var u model.User
	var r *gorm.DB

	if len(in.Account) > 0 {
		r = l.svcCtx.DB.Where("account = ? ", in.Account).First(&u)
	} else {
		r = l.svcCtx.DB.First(&u, in.Id)
	}

	if r.Error != nil {
		return &user.UserInfoRsp{
			Code:    1,
			Message: tool.ErrorCode[1],
		}, nil
	}

	return &user.UserInfoRsp{
		Code: 0,
		User: &user.UserDetail{
			Id:       uint64(u.ID),
			Account:  u.Account,
			Name:     u.Name,
			Avatar:   u.Avatar,
			Status:   u.Status,
			Password: u.Password,
		},
	}, nil
}
