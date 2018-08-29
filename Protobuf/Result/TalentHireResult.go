package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//雇佣经理返回 1177
func TalentHireResult(Uid int32,Data []byte) *AutoMsg.TalentHireResult{
	TalentHireResult := &AutoMsg.TalentHireResult{}
	proto.Unmarshal(Data,TalentHireResult)
	Param,_ := json.Marshal(TalentHireResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1177,
		Name:  "雇佣经理返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalentHireResult
}
