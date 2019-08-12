package main

import (
	"encoding/json"
	"fmt"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/cache"
)

func main() {
	channelMsg := ScoreChannelMsg{UserId: 1, Score: 1001}
	channelMsg.MsgType = 1
	channelMsg.Channel = "registerEmail"
	Publish(&channelMsg)
}

func Publish(processor Processor) {

	fmt.Println("channel name:", processor.GetChannel())

	jsonBytes, err := json.Marshal(processor)
	if err != nil {
		panic(err)
	}

	msgString := string(jsonBytes)

	err = cache.GetCache().Publish(processor.GetChannel(), msgString).Err()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
