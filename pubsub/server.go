package main

import (
	"encoding/json"
	"fmt"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/cache"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	Subscribe()
}

func Subscribe() {
	fmt.Println("Staring subscribe score....")
	pubSub := cache.GetCache().Subscribe("score", "registerEmail")
	defer pubSub.Close()

	var channel string
	var payload string

	for {

		msg, err := pubSub.Receive()
		if err != nil {
			panic(err)
		}

		// channel = ""
		// payload = ""
		fmt.Println(msg)

		switch m := msg.(type) {

		case *redis.Subscription:
			log.Printf("Subscription: %s, %d, %s", m.Channel, m.Count, m.Kind)
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
			fmt.Println(channel, payload)
			go run(channel, payload)

		default:
			log.Printf("Default: %s", m)
		}

	}
}

func run(channel, payload string) {
	switch channel {
	case "score":
		fmt.Println("开始送积分！！！")
		scoreChannelMsg := ScoreChannelMsg{}
		json.Unmarshal([]byte(payload), &scoreChannelMsg)
		scoreChannelMsg.Run()
	case "registerEmail":
		fmt.Println("注册发送邮件！！！")
	default:
		log.Printf("run: %s %s ", channel, payload)
	}

}
