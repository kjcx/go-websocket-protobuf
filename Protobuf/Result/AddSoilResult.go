package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//庄园购买地块返回 1110
func AddSoilResult(Uid int32,Data []byte) *AutoMsg.AddSoilResult {
	AddSoilResult := &AutoMsg.AddSoilResult{}
	proto.Unmarshal(Data,AddSoilResult)
	Param,_ := json.Marshal(AddSoilResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1110,
		Name:  "庄园购买地块返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return AddSoilResult
}
