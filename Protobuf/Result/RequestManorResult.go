package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//1082 庄园请求返回
func RequestManorResult(Uid int32,Data []byte) *AutoMsg.RequestManorResult{
	RequestManorResult := &AutoMsg.RequestManorResult{}
	proto.Unmarshal(Data,RequestManorResult)
	Param,_ := json.Marshal(RequestManorResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1082,
		Name:  "庄园请求返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return RequestManorResult
}
