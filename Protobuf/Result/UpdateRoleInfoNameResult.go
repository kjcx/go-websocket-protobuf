package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//修改角色名称返回 1142
func UpdateRoleInfoNameResult(Uid int32,Data []byte) *AutoMsg.UpdateRoleInfoNameResult {
	UpdateNameResult := &AutoMsg.UpdateRoleInfoNameResult{}
	proto.Unmarshal(Data,UpdateNameResult)
	Param,_ := json.Marshal(UpdateRoleInfoIconResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1141,
		Name:"更改角色昵称返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return UpdateNameResult
}
