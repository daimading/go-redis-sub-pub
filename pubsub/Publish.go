package main

import (
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/cache"
	"score_sub/pubsub/message"
	"github.com/json-iterator/go"
)

func main() {
	channelMsg := message.ChannelMsgExample{UserId: 1, Score: 1001}
	channelMsg.MsgType = 1
	channelMsg.Channel = CHANNEL_NAME_EXAMPLE_1
	Publish(&channelMsg)
}

func Publish(processor message.Processor) {

	msgBytes, err := jsoniter.Marshal(processor)
	if err != nil {
		panic(err)
	}

	err = cache.GetCache().Publish(processor.GetChannel(), string(msgBytes)).Err()

	if err != nil {
		panic(err)
	}
}
