/*
2-operation_redis.go: go 操作 Redis

	go get github.com/gomodule/redigo/redis
*/
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func operationRedis() {
	c, err := redis.Dial("tcp", "192.168.100.115:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")

	defer c.Close()

	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
