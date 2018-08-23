package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

//解析角色 1060
func CreateRoleResult(Data []byte) *AutoMsg.CreateRoleResult {
	fmt.Println("解析角色")
	CreateRole := &AutoMsg.CreateRoleResult{}
	err := proto.Unmarshal(Data, CreateRole)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(CreateRole.GetRoleId())
	return CreateRole

}