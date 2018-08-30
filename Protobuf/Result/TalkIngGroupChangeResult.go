package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//谈判团换任返回 1158
func TalkIngGroupChangeResult(Uid int32,Data []byte) *AutoMsg.TalkIngGroupChangeResult {
	TalkIngGroupChangeResult := &AutoMsg.TalkIngGroupChangeResult{}
	proto.Unmarshal(Data,TalkIngGroupChangeResult)
	Param,_ := json.Marshal(TalkIngGroupChangeResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1158,
		Name:  "谈判团换任返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalkIngGroupChangeResult
}
