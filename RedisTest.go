package main


/*
import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	fmt.Println("1"+val2+"1")
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	err=client.Set("id","name",time.Duration(10*time.Second)).Err()
	time.Sleep(9*time.Second)
	val3,err:=client.Get("id").Result()
	 if err==redis.Nil {
		 fmt.Println("NO id")
	 }else	if err!=nil {
			 panic(err)
		 } else {
		 fmt.Println(val3)
	}
	_, err = client.HSet("ming1", "id", "12313").Result()
	key4, err := client.HGet("ming1", "id").Result()
	if err==redis.Nil{
		fmt.Println("nb")
	}else if err!=nil{
		panic(err)
	}else{
		fmt.Println(key4)
	}
}*/