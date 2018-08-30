package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//更改居民返回 2034
func TalkIngGroupResidentChangeResult(Uid int32,Data []byte) *AutoMsg.TalkIngGroupResidentChangeResult {
	TalkIngGroupResidentChangeResult := &AutoMsg.TalkIngGroupResidentChangeResult{}
	proto.Unmarshal(Data,TalkIngGroupResidentChangeResult)
	Param,_ := json.Marshal(TalkIngGroupResidentChangeResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2034,
		Name:  "更改居民返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalkIngGroupResidentChangeResult
}
