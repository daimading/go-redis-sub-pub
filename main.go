package main

import (
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/database"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/log"
	"github.com/jinzhu/gorm"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/http_server/server"
	"fmt"
	"git.dev.tencent.com/yctc/golang-basic-framework.git/core/cache"
)

func main() {
	redisClient := cache.GetCache()
	defer redisClient.Close()

	// 指定通道
	pubSub := redisClient.Subscribe("mychannel1")
	defer pubSub.Close()

	go func() {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}()

	// channel 通道  message 放入通道中订阅的信息
	err := redisClient.Publish("mychannel1", "hello word ！！！").Err()
	if err != nil {
		panic(err)
	}
}

// 订阅的信息
func giveScore(assert string) (err error) {

	db := database.GetDB()
	//  查询用户积分资产信息
	err = db.First(&assert).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Error(log.NewCommon(map[string]interface{}{"message": "database error!"}, err))
		return server.ErrorFromCode(server.STATUS_SERVER_ERROR_CODE)
	} else if gorm.IsRecordNotFoundError(err) {
		// 为老用户创建积分资产表
	} else {
		// 修改用户积分资产，新积分+要添加的积分
		err = database.GetDB().Updates("").Error
		if err != nil {
			log.Error(log.NewCommon(map[string]interface{}{"message": "update assets error!"}, err))
			return server.ErrorFromCode(server.STATUS_SERVER_ERROR_CODE)
		}
	}

	// 向mongodb中添加积分记录(过期时间)
	if err != nil {
		log.Error(log.NewCommon(map[string]interface{}{"message": "add score error!"}, err))
		return server.ErrorFromCode(server.STATUS_SERVER_ERROR_CODE)
	}
	return
}
