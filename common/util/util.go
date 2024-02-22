package util

import (
	"fmt"
)

var Util = util{}

type util struct {
}

func (u *util) ConvertInt64ToStrs(list []int64) []string {
	if list == nil || len(list) == 0 {
		return nil
	}

	strs := make([]string, len(list))
	for i, v := range list {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return strs
}
