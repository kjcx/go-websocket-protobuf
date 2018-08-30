package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//培训员工返回
func CultivateEmployeeResult(Uid int32,Data []byte) *AutoMsg.CultivateEmployeeResult{
	CultivateEmployeeResult := &AutoMsg.CultivateEmployeeResult{}
	proto.Unmarshal(Data,CultivateEmployeeResult)
	Param,_ := json.Marshal(CultivateEmployeeResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1121,
		Name:  "培训员工返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return CultivateEmployeeResult
}
