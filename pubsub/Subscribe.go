package main

import (
	"encoding/json"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/cache"
	"score_sub/pubsub/message"
	"github.com/sirupsen/logrus"
)

const (
	CHANNEL_NAME_EXAMPLE_1 = "channel_name_1"
	CHANNEL_NAME_EXAMPLE_2 = "channel_name_2"
)

func main() {
	Subscribe()
}

func Subscribe() {
	logrus.Info("Staring subscribe ....")
	pubSub := cache.GetCache().Subscribe(CHANNEL_NAME_EXAMPLE_1, CHANNEL_NAME_EXAMPLE_2)
	defer pubSub.Close()

	ch := pubSub.Channel()

	for {
		msg, ok := <-ch
		if !ok {
			break
		}

		switch msg.Channel {

		case CHANNEL_NAME_EXAMPLE_1:
			go run(msg.Channel, msg.Payload)
		case CHANNEL_NAME_EXAMPLE_2:
			go run(msg.Channel, msg.Payload)
		default:
			logrus.Info("Default: %s", msg)
		}
	}
}

func run(channel, payload string) {

	channelMsgExample := message.ChannelMsgExample{}
	json.Unmarshal([]byte(payload), &channelMsgExample)

	switch channel {

	case CHANNEL_NAME_EXAMPLE_1:
		channelMsgExample.Run()
	case CHANNEL_NAME_EXAMPLE_2:

	default:
		logrus.Info("run: %s %s ", channel, payload)
	}
}
