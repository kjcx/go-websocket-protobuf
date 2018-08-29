package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//解雇经理返回
func TalentFireResult(Uid int32,Data []byte) *AutoMsg.TalentFireResult{
	TalentFireResult := &AutoMsg.TalentFireResult{}
	proto.Unmarshal(Data,TalentFireResult)
	Param,_ := json.Marshal(TalentFireResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1178,
		Name:  "解雇经理返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalentFireResult
}
