package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)


//签到列表 1168
func SignResult(Uid int32,Data []byte) *AutoMsg.SignResult  {
	SignResult := &AutoMsg.SignResult{}
	err := proto.Unmarshal(Data, SignResult)
	if err != nil {
		fmt.Println("解析出错")
	}
	Param,_ := json.Marshal(SignResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1168,
		Name:"签到列表",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return SignResult
}
