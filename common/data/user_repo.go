package data

import (
	"context"
	"github.com/im/common/api"
	"github.com/im/common/data/ent"
	"github.com/im/common/data/ent/userinfo"
)

var UserRepo = &userRepo{}

type userRepo struct{}

func (p *userRepo) FindByUsernameAndPassword(ctx context.Context, username string, password string) *ent.UserInfo {
	return DataM.GetDBClient().UserInfo.Query().Where(userinfo.UsernameEQ(username), userinfo.PasswordEQ(password)).FirstX(ctx)
}

func (p *userRepo) CountByUsername(ctx context.Context, username string) int {
	return DataM.GetDBClient().UserInfo.Query().Where(userinfo.UsernameEQ(username)).CountX(ctx)
}

func (p *userRepo) Create(ctx context.Context, req *api.UserRegisterReq) *ent.UserInfo {
	return DataM.GetDBClient().UserInfo.Create().SetUsername(req.Username).SetPassword(req.Password).SetNickname(req.Nickname).SetAvatar(req.Avatar).SetStatus(0).SaveX(ctx)
}
