package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//任务返回 1029
func MissionDataResult(Uid int32,Data []byte) *AutoMsg.MissionDataResult {
	MissionDataResult := &AutoMsg.MissionDataResult{}
	proto.Unmarshal(Data,MissionDataResult)
	Param,_ := json.Marshal(MissionDataResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1029,
		Name:  "任务返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return MissionDataResult
}
