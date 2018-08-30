package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//批量删除邮件返回 1113
func DelMailsResult(Uid int32,Data []byte) *AutoMsg.DelMailsResult{
	DelMailsResult := &AutoMsg.DelMailsResult{}
	proto.Unmarshal(Data,DelMailsResult)
	Param,_ := json.Marshal(DelMailsResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1113,
		Name:  "批量删除邮件返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return DelMailsResult
}
