package main

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
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

	var lines []string

	for i := 0; true; i += len(lines) {
		lines, err = c.Cmd("lrange", "queue1", i, -1).List()
		handleErr(err)

		if len(lines) > 0 {
			fmt.Println(lines)
		}
	}

}
