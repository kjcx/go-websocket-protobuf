package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//随机请求庄园返回 2014
func RandManorResult(Uid int32,Data []byte) *AutoMsg.RandManorResult {
	RandManorResult := &AutoMsg.RandManorResult{}
	proto.Unmarshal(Data,RandManorResult)
	Param,_ := json.Marshal(RandManorResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2014,
		Name:  "随机请求庄园返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return RandManorResult
}
