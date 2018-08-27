package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//礼物列表 2056
func GiveListResult(Uid int32,Data []byte) *AutoMsg.GiveListResult {
	GiveListResult := &AutoMsg.GiveListResult{}
	proto.Unmarshal(Data,GiveListResult)
	Param,_ := json.Marshal(GiveListResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2056,
		Name:  "礼物列表",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return GiveListResult
}
