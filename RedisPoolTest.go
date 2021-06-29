package main
/*
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool


func init(){
	pool=&redis.Pool{
		MaxIdle: 20,
		MaxActive: 0,
		IdleTimeout: 0,
		Dial:func()(redis.Conn,error){
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}
}
func main(){
	c:=pool.Get()
	defer c.Close()
	_,err:=c.Do("set","abc",100)
	if err!=nil{
		fmt.Println("Do err",err)
		return
	}

	r,err:=redis.Int(c.Do("get","abc"))
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(r)
	fmt.Printf("%T",r)
	pool.Close()
}
*/