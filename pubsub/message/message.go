package message

import (
	"github.com/sirupsen/logrus"
)

// BaseModel
type ChannelMsg struct {
	Channel string
	MsgType int
}

type Processor interface {
	Run() interface{}
	GetChannel() string
}

type ChannelMsgExample struct {
	ChannelMsg
	UserId     int32
	Score      int32
}

// GetChannel get channel
func (cm *ChannelMsg) GetChannel() string {
	return cm.Channel
}

// Run
func (cm *ChannelMsgExample) Run() interface{} {
	logrus.Info("ChannelMsg...", cm.Score, cm.UserId)
	logrus.Println("message operation")
	return true
}
