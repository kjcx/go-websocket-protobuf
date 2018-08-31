package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//偷菜返回 1091
func StealSemenResult(Uid int32,Data []byte) *AutoMsg.StealSemenResult {
	StealSemenResult := &AutoMsg.StealSemenResult{}
	proto.Unmarshal(Data,StealSemenResult)
	Param,_ := json.Marshal(StealSemenResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1091,
		Name:  "偷菜返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return StealSemenResult
}
