package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//商战请求返回 2044
func BusinessPkResult(Uid int32,Data []byte) *AutoMsg.BusinessPkResult {
	BusinessPkResult := &AutoMsg.BusinessPkResult{}
	proto.Unmarshal(Data,BusinessPkResult)
	Param,_ := json.Marshal(BusinessPkResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2044,
		Name:  "商战请求返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return BusinessPkResult
}
