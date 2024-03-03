package data

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/im/common/data/ent"
	"github.com/im/common/data/ent/chatlist"
)

var ChatListRepo = &chatListRepo{}

type chatListRepo struct {
}

func (p *chatListRepo) CreateOrUpdate(ctx context.Context, datas ...*ent.ChatList) {
	for _, chat := range datas {
		count := DataM.GetDBClient().ChatList.Update().
			SetMaxMsgID(chat.MaxMsgID).
			SetUts(chat.Uts).
			SetType(chat.Type).Where(chatlist.UIDEQ(chat.UID), chatlist.ChannelEQ(chat.Channel), chatlist.ChatIDEQ(chat.ChatID)).SaveX(ctx)
		if count <= 0 {
			_, err := DataM.GetDBClient().ChatList.Create().
				SetUID(chat.UID).
				SetChannel(chat.Channel).
				SetChatID(chat.ChatID).
				SetMaxMsgID(chat.MaxMsgID).
				SetUts(chat.Uts).
				SetType(chat.Type).Save(ctx)
			if err != nil {
				logger.Errorf("ChatListRepo CreateOrUpdate failed:%v", err)
			}
		}

	}
}
