package main

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"net/http"
	"io/ioutil"
	"net/url"
)

func main() {
	//var i int

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.119:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	fmt.Println(time.Now())
	//for i=1;i<100000;i++ {
	//	key := strconv.Itoa(i)
		//Token := HttpController.GetToken("192.168.31.232:9501",key)
		//fmt.Println(Token)
		//Redis.InsertToken(key,Token,key)
	//	err := client.Set(key,"123",0).Err()
	//	//go Redis.InsertToken(key,"Token",key)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fmt.Println(time.Now())
	//for i=1;i<10000;i++ {
	//	go geturl()
	//
	//}
	posturl()

	time.Sleep(time.Second*1)
}
func geturl(){
	resp,err :=http.Get("https://www.ka-cn.com/spgoods/goods_zfb.php?id=1757");
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	time.Sleep(time.Second*1)
}

func posturl(){
	postValue := url.Values{
		"email": {"xx@xx.com"},
		"password": {"123456"},
		"is_ajax_request": {"1"},
		"agreement": {"1"},
		"register_types": {"1"},
		"direct_shopping_key":{"1"},
		"order_id":{"111"},
		"confirm_password": {"111111"},
		"back_act": {"https%3A%2F%2Fwww.ka-cn.com%2Farticle-1402.html"},

	}

	resp, err := http.PostForm("http://www.ka-cn.com/user.php?act=register", postValue)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
