package api

import (
	"github.com/im/common/constants"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ConvertPBResultToJsonResult(result *Result) (*JsonResult, error) {
	var data interface{}
	if result.Code == constants.SuccessCode {
		data1, err := result.Data.UnmarshalNew()
		if err != nil {
			return nil, err
		}
		data = data1
	}

	return &JsonResult{
		Code: int(result.Code),
		Msg:  result.Msg,
		Data: data,
	}, nil
}
