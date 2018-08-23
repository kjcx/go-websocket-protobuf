package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)
//修改角色名称返回 1142
func UpdateRoleInfoNameResult(Data []byte) *AutoMsg.UpdateRoleInfoNameResult {
	UpdateNameResult := &AutoMsg.UpdateRoleInfoNameResult{}
	proto.Unmarshal(Data,UpdateNameResult)
	return UpdateNameResult
}
