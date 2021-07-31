package logic

import (
	"context"
	"gitee.com/fireflylove/user-svc/model"
	"gitee.com/fireflylove/user-svc/tool"

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

	var err error
	var r *model.User

	if len(in.Account) > 0 {
		r, err = l.svcCtx.UserModel.FindOneByAccount(in.Account)
	} else {
		r, err = l.svcCtx.UserModel.FindOne(in.Id)
	}

	if err != nil {
		return &user.UserInfoRsp{
			Code:    1,
			Message: tool.ErrorCode[1],
		}, nil
	}

	return &user.UserInfoRsp{
		Code: 0,
		User: &user.UserDetail{
			Id:       r.Id,
			Account:  r.Account,
			Name:     r.Name,
			Avatar:   r.Avatar,
			Status:   r.Status,
			Password: r.Password,
		},
	}, nil
}
