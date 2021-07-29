package logic

import (
	"context"

	"user-svc/internal/svc"
	"user-svc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user.UserListReq) (*user.UserListRsp, error) {

	m := l.svcCtx.UserModel

	r, err := m.FindByIds(in.Id)

	if err != nil {
		return &user.UserListRsp{
			Code: 0,
		}, nil
	}

	var list []*user.UserInfo

	for _, item := range *r {
		list = append(list, &user.UserInfo{
			Id:      item.Id,
			Account: item.Account,
			Name:    item.Name,
			Avatar:  item.Avatar,
			Status:  item.Status,
		})
	}

	return &user.UserListRsp{
		Code: 0,
		User: list,
	}, nil
}
