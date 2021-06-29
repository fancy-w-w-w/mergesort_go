package redismq

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
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

func BatchPushQueue(queueName string, keys []int64) (err error) {
	con:= pool.Get()
	defer con.Close()
	_, err = con.Do("lpush", redis.Args{}.Add(queueName).AddFlat(keys)...)
	return
}
//timeout is seconds which command 'brpop' will block when queue is empty.
//brpop会一致阻塞住直到队列中有元素，但是它支持设置timeout，当阻塞时间超过timeout时，pop会返回nil。当timeout设置为0时，表示阻塞时间无限制。
//brpop支持监听多个list，因此它有两个返回值，第一个返回值是list的名称，即key的名称，第二个返回值是pop出来的元素。我们只是监听一个list，因此我们会取返回值中的第二个元素。
func PopQueue(queueName string, timeout int) (data int64, err error) {
	con := pool.Get()
	defer con.Close()
	nameAndData, err := redis.Strings(con.Do("brpop", queueName, timeout))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	if len(nameAndData) > 1 {
		str := nameAndData[1]
		data, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetQueueLen(queueName string)(len int,err error){
	con:=pool.Get()
	defer con.Close()
	len,err=redis.Int(con.Do("llen",queueName))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			return
		}
		return
	}
	return

}


