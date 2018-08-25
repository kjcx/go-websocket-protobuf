package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"encoding/json"
	"time"
)

//改变时装请求返回 1055
func ChangeAvatarResult(Uid int32,Data []byte) *AutoMsg.ChangeAvatarResult {
	ChangeAvatarResult := &AutoMsg.ChangeAvatarResult{}
	proto.Unmarshal(Data,ChangeAvatarResult)
	Param,_ := json.Marshal(ChangeAvatarResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1055,
		Name:"改变时装请求返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)
	return ChangeAvatarResult
}
