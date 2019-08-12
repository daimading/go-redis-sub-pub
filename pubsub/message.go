package main

import "fmt"

type ChannelMsg struct {
	Channel string
	MsgType int
}

func (cm *ChannelMsg) GetChannel() string {
	return cm.Channel
}

// base
type Processor interface {
	Run() interface{}
	GetChannel() string
}

type ScoreChannelMsg struct {
	ChannelMsg
	UserId int32
	Score int32
}

func (cm *ScoreChannelMsg) Run() interface{}{
	fmt.Println("ChannelMsg...", cm.Score, cm.UserId)
	fmt.Println("积分操作具体详情")
	return true
}
