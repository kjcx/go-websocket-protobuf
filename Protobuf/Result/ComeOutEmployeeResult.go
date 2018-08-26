package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//调入调出员工返回 1126
func ComeOutEmployeeResult(Uid int32,Data []byte) *AutoMsg.ComeOutEmployeeResult{
	ComeOutEmployeeResult := &AutoMsg.ComeOutEmployeeResult{}
	proto.Unmarshal(Data,ComeOutEmployeeResult)
	Param,_ := json.Marshal(ComeOutEmployeeResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1126,
		Name:"调入调出员工返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return ComeOutEmployeeResult
}
