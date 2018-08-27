package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	str := `{"RoleId":13,"ItemCount":{"501":13}}`
	type T struct {
		RoleId int32
		ItemCount map[int32]int64
	}
	A := &T{}
	json.Unmarshal([]byte(str),A)
	fmt.Println(A.ItemCount)
}