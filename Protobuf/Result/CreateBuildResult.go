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
func CreateBuildResult(Uid int32,Data []byte) *AutoMsg.CreateBuildResult {
	CreateBuild := &AutoMsg.CreateBuildResult{}
	proto.Unmarshal(Data,CreateBuild)
	fmt.Println(CreateBuild)
	Param,_ := json.Marshal(CreateBuild)
	fmt.Println("建造店铺返回")
	fmt.Println(Param)
	log := &Mgo.Log{
		Uid:   Uid,
		MsgId: 1058,
		Name:  "建造店铺返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:"",
	}
	log.InsertLog()
	return CreateBuild
}
