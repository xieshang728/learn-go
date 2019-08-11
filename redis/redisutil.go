package redisutil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gpmgo/gopm/modules/log"
)

var redisPool *redis.Pool

const (
	SET   = "SET"
	SETNX = "SETNX"
	//GET
	GET   = "GET"
	DEL   ="DEL"
	EXPIRE = "EX"
	EMPTY = ""
	ZERO = 0
)

func init() {
	redisPool =  &redis.Pool{
		MaxIdle:3,
		MaxActive:5,
		Dial: func() (conn redis.Conn, e error) {
			c,err := redis.Dial("tcp","39.96.5.21:6379")
			if err != nil {
				return nil,err
			}
			return c,err
		},
	}
}


func Get(key string) string{
	log.Info("redisutil.get.operation.key: %s",key)
	conn := redisPool.Get()
	result,err := redis.String(conn.Do(GET,key))
	if err != nil{
		log.Error("resultutil.set err: %s",err)
		return EMPTY
	}
	log.Info("redisutil.get.operation.value: %s",result)
	return result
}


func Set(key string, value string){
	log.Info("redisutil.set.operation.key: %s value:%s",key,value)
	conn := redisPool.Get()
	result,err := conn.Do(SET,key,value)
	if err != nil{
		log.Error("redisutil.set.operation.err: %s",err.Error())
		return
	}
	log.Info("redisutil.set.operation.result: %s",result)
}


func Del(key string){
	log.Info("redisutil.del.operation.key: %s",key)
	conn := redisPool.Get()
	result,err := conn.Do(DEL,key)
	if err != nil {
		log.Error("redisutil.del.operation.err: %s",err.Error())
		return
	}
	log.Info("redisutil.del.operation.result: %s",result)
}

func SetForExpire(key string, value string, expire int){
	log.Info("redisutil.set.expire.operation.key: %s,value: %s,expire %d",key,value,expire)
	conn := redisPool.Get()
	result,err := conn.Do(SET,key,value,EXPIRE,expire)
	if err != nil {
		log.Error("redisutil.set.expire.operation.err %s",err.Error())
		return
	}
	log.Info("redisutil.set.expire.operation.result %s",result)

}

func SetNX(key string,value string) int{
	log.Info("redisutil.set.nx.operation.key: %s,value:%s",key,value)
	conn := redisPool.Get()
	result,err := redis.Int(conn.Do(SETNX,key,value))
	if err != nil{
		log.Error("redisutil.setnx.operation.err: %s",err.Error())
		return ZERO
	}
	log.Info("redisutil.setnx.operation.result.%d",result)
	return result
}

