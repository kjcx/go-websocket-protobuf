package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)
//创建店铺结构
type CreateBuild struct {
	Pos int32
	AeraId int32
	ShopType int32
}
//发送创建店铺 1005
func CreateBuildReq(CreateBuild *CreateBuild)[]byte {
	CreateBuildReq := &AutoMsg.CreateBuildReq{
		Pos:CreateBuild.Pos,
		AreaId:CreateBuild.AeraId,
		ShopType:CreateBuild.ShopType,
	}
	Data,_ := proto.Marshal(CreateBuildReq)
	return Send(1005,Data)

}
