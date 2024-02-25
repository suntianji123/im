package util

import (
	"fmt"
	"github.com/im/common/conf"
)

var ChannelUtil = &channelUtil{
	mp: make(map[string]*channelCfgItem),
}

type channelCfgItem struct {
	chatType int
	channel  int
	appIds   []int
}

type channelUtil struct {
	mp map[string]*channelCfgItem
}

func (p *channelUtil) Init(cfg *conf.ChannelConfig) {
	items := p.buildItems(cfg.Items)
	for _, item := range items {
		p.mp[p.genKeyByItem(item)] = item
	}
}

func (p *channelUtil) buildItems(cfgs []*conf.ChannelConfigItem) []*channelCfgItem {
	result := make([]*channelCfgItem, 0)
	for _, cfg := range cfgs {

		for _, item := range cfg.Config {
			result = append(result, &channelCfgItem{
				chatType: cfg.ChatType,
				channel:  item.Channel,
				appIds:   item.AppIds,
			})
		}
	}
	return result
}

func (p *channelUtil) genKeyByItem(item *channelCfgItem) string {
	return p.genKey(item.chatType, item.channel)
}

func (p *channelUtil) genKey(chatType int, channel int) string {
	return fmt.Sprintf("%d:%d", chatType, channel)
}

func (p *channelUtil) GetAppIds(chatType int, channel int, appId int) []int {
	key := p.genKey(chatType, channel)
	result := make([]int, 0)
	exists := false
	if v, ok := p.mp[key]; ok {
		for _, v1 := range v.appIds {
			if v1 == appId {
				exists = true
				break
			}
		}
		result = append(result, v.appIds...)
	}

	if !exists {
		result = append(result, appId)
	}
	return result
}
