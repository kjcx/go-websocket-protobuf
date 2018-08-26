package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//解析角色 1060
func CreateRoleResult(Uid int32,Data []byte) *AutoMsg.CreateRoleResult {
	fmt.Println("解析角色")
	CreateRole := &AutoMsg.CreateRoleResult{}
	err := proto.Unmarshal(Data, CreateRole)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(CreateRole.GetRoleId())
	Param,_ := json.Marshal(CreateRole)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1060,
		Name:"解析角色",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return CreateRole

}