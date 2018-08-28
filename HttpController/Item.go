package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"encoding/json"
	"fmt"
)

type ItemCount struct {
	ItemId int32
	Count int64
}
type ItemCounts struct {
	Items []ItemCount
}

//添加道具
func AddItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data:= &ItemCounts{}
	json.Unmarshal(param.GetParam().Body,data)
	fmt.Println(data)
	for k,v := range data.Items{
		Common.AddItemSend(param.GetParam().Uid,v.ItemId,v.Count)
		fmt.Println(k)
		fmt.Println(v)
	}
}