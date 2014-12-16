package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
	"time"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:3003")
	handleErr(err)
	defer c.Close()

	c.Cmd("del", "queue1")

	for i := 0; true; i++ {
		fmt.Println("Publishing:", i)
		c.Cmd("rpush", "queue1", i)
		time.Sleep(time.Duration((1000)) * time.Millisecond)
	}
}
