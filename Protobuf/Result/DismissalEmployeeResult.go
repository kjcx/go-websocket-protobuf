package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//解雇员工返回 1122
func DismissalEmployeeResult(Uid int32,Data []byte) *AutoMsg.DismissalEmployeeResult  {
	DismissalEmployeeResult := &AutoMsg.DismissalEmployeeResult{}
	proto.Unmarshal(Data,DismissalEmployeeResult)
	Param,_ := json.Marshal(DismissalEmployeeResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1122,
		Name:  "解雇员工返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return DismissalEmployeeResult
}
