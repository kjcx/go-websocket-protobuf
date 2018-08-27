package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//1176 人才市场列表返回
func GetTalentListResult(Uid int32,Data []byte) *AutoMsg.GetTalentListResult{
	GetTalentListResult := &AutoMsg.GetTalentListResult{}
	proto.Unmarshal(Data,GetTalentListResult)
	Param,_ := json.Marshal(GetTalentListResult)
	log := &Mgo.Log{
		Uid: Uid,
		MsgId: 1176,
		Name: "人才市场列表返回",
		Param: string(Param),
		Date: time.Now(),
		Msg: "无",
	}
	log.InsertLog()
	return GetTalentListResult
}
