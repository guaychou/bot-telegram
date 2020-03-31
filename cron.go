package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gomodule/redigo/redis"
	rdc "github.com/guaychou/redisClient"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var grupchatid,err=strconv.ParseInt(os.Getenv("GRUP_CHAT"), 10, 64)
var grupmessage = tgbotapi.NewMessage(grupchatid, "")

func cronfunction (c *cron.Cron, conn *redis.Conn, bot *tgbotapi.BotAPI){
	log.Info("Cron Job Started . . . ")
	c.AddFunc("*/1 * * * *", func() { redisHealthCheck(conn)})
	c.AddFunc("0 0 0,12,18 ? * * *", func() { sendMessage(bot) })
	c.Start()
}

func redisHealthCheck(conn *redis.Conn){
	log.Info("Healthcheck Status: "+ rdc.RedisClientPing(*conn))
}

func sendMessage(bot *tgbotapi.BotAPI){
	grupmessage.Text=corona("/corona indonesia")
	_,err:=bot.Send(grupmessage)
	if err!=nil{
		log.Error(err)
	}else{
		log.Info("Message Status: Delivered")
	}
}