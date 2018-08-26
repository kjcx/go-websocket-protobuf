package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//1117 招聘返回
func RefStaffResult(Uid int32,Data []byte) *AutoMsg.RefStaffResult{
	RefStaffResult := &AutoMsg.RefStaffResult{}
	proto.Unmarshal(Data,RefStaffResult)
	Param,_ := json.Marshal(RefStaffResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1117,
		Name:"招聘返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return  RefStaffResult
}