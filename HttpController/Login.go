package HttpController

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

func GetToken(httpurl string ,key string) string{
	time.Sleep(time.Microsecond * 10000)
	url := fmt.Sprintf("http://%s%s%s", httpurl, "/Account/CKlogin?key=", key)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	type Token struct {
		Token string
	}

	type Result struct {
		Data *Token
		Msg  string
	}

	//fmt.Println(string(body))

	L := Result{}
	json.Unmarshal([]byte(body), &L)
	fmt.Println(L.Data.Token)
	return L.Data.Token

}
