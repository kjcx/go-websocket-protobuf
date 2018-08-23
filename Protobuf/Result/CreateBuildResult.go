package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Mgo"
	"encoding/json"
	"time"
	"fmt"
)
//建造店铺返回 1058
func CreateBuildResult(Data []byte) *AutoMsg.CreateBuildResult {
	CreateBuild := &AutoMsg.CreateBuildResult{}
	proto.Unmarshal(Data,CreateBuild)
	fmt.Println(CreateBuild)
	mongo := Mgo.Mongo()
	Param,_ := json.Marshal(CreateBuild)
	fmt.Println("建造店铺返回")
	fmt.Println(Param)
	LogData := &Mgo.Log{
		Uid:   14,
		MsgId: 1058,
		Name:  "建造店铺返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:"",
	}
	Mgo.InsertLog(mongo, LogData)
	return CreateBuild
}
