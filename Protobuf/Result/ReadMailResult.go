package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//阅读邮件返回 1114
func ReadMailResult(Uid int32,Data []byte) *AutoMsg.ReadMailResult{
	ReadMailResult := &AutoMsg.ReadMailResult{}
	proto.Unmarshal(Data,ReadMailResult)
	Param,_ := json.Marshal(ReadMailResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1114,
		Name:  "阅读邮件返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return ReadMailResult
}
