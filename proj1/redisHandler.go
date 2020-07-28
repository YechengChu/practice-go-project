package main

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

var c redis.Conn

func errCheck(info string, err error) {
	if err != nil {
		fmt.Printf("%s %v\n", info, err)
		os.Exit(-1)
	}
}

func connectRedis() {
	var err error
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	errCheck("Connect to redis error:", err)
}

func closeRedis() {
	c.Close()
}

func addInRedis(sessionIDGiven string, account string) {
	_, err := c.Do("SET", sessionIDGiven, account)
	errCheck("redis set failed:", err)
}

func isLoggedInRedis(sessionIDGiven string) bool {
	is_key_exit, err := redis.Bool(c.Do("EXISTS", sessionIDGiven))
	errCheck("error:", err)
	return is_key_exit
}

func getInRedis(sessionIDGiven string) string {
	account, err := redis.String(c.Do("GET", sessionIDGiven))
	errCheck("redis get failed:", err)
	return account
}

func removedInRedis(sessionIDGiven string) {
	_, err := c.Do("DEL", sessionIDGiven)
	errCheck("redis delete failed:", err)
}
