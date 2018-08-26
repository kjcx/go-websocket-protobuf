package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//加入游戏返回 1066
func JoinGameResult(Uid int32,Data []byte) *AutoMsg.JoinGameResult {
	JoinGameResult := &AutoMsg.JoinGameResult{}
	err := proto.Unmarshal(Data, JoinGameResult)
	if err != nil {
		fmt.Println("marshaling error:1012 ", err.Error())
	}
	Param,_ := json.Marshal(JoinGameResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1066,
		Name:"进入游戏返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return JoinGameResult
}
