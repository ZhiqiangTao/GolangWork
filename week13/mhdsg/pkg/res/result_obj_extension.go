package res

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Success(data proto.Message) *ResultObj {
	resObj := &ResultObj{}
	resObj.Code = 200
	resObj.Msg = "ok"

	any, _ := anypb.New(data)
	resObj.Data = any
	return resObj
}

func Fail(code int32, msg string, data string) *ResultObjString {
	resObj := &ResultObjString{}
	resObj.Code = code
	resObj.Msg = msg
	resObj.Data = data
	return resObj
}
