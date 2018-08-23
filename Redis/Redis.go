package Redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"encoding/json"
)

func InsertToken(Key string ,Token string,Id string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.119:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	type T struct {
		Key string
		Token string
		Id string
	}
	var str = &T{Key,Token,Id}

	b, err := json.Marshal(str)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}


	fmt.Println("json:", string(b))
	var SetKey = "NewToken:"
	s := SetKey + Key
	err = client.Set(s,b , 0).Err()
	if err != nil {
		panic(err)
	}

	//val, err := client.Get("key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)

	//val2, err := client.Get("key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exists")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
}

func Insert(value []byte){
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.119:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	client.Set("JoinGame",value,0).Err()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

}
