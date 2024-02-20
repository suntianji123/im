package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/api"
	"github.com/im/common/data/ent"
	"github.com/im/common/data/ent/friend"
	"github.com/im/common/data/ent/userinfo"
	"strconv"
)

var FriendRepo = friendRepo{}

type friendRepo struct{}

func (*friendRepo) ConvertEntUserInfoToPbUserInfo(user *ent.UserInfo) *api.UserInfo {
	return &api.UserInfo{
		Id:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Ext:      user.Ext,
	}
}

func (*friendRepo) FindFriends(ctx context.Context, req *api.FriendListReq) ([]*ent.UserInfo, error) {
	idStrs := Data.Db.Friend.Query().Where(friend.UIDEQ(req.Uid), friend.IDGT(req.MinFriendId)).Limit(int(req.Size)).Select(friend.FieldPeerUID).StringsX(ctx)
	if idStrs == nil || len(idStrs) == 0 {
		return nil, nil
	}

	var err error
	ids := make([]int64, len(idStrs))
	for i, v := range idStrs {
		ids[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			logger.Errorf("FriendRepo FindFriends failed:%v", err)
			return nil, err
		}
	}
	return Data.Db.UserInfo.Query().Where(userinfo.IDIn(ids...)).AllX(ctx), nil
}
