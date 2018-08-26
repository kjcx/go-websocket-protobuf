package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)
//更改头像返回1141
func UpdateRoleInfoIconResult(Uid int32,Data []byte) *AutoMsg.UpdateRoleInfoIconResult{
	UpdateRoleInfoIconResult := &AutoMsg.UpdateRoleInfoIconResult{}
	proto.Unmarshal(Data,UpdateRoleInfoIconResult)
	Param,_ := json.Marshal(UpdateRoleInfoIconResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1141,
		Name:"更改头像返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return  UpdateRoleInfoIconResult
}
