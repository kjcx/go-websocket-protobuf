package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)

// 刷新商城商品返回 1109
func RefDropShopResult(Data []byte) *AutoMsg.RefDropShopResult{
	RefDropShopResult := &AutoMsg.RefDropShopResult{}
	proto.Unmarshal(Data,RefDropShopResult)
	return RefDropShopResult
}
