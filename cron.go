package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gomodule/redigo/redis"
	rdc "github.com/guaychou/redisClient"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

var grupchatid,err=strconv.ParseInt(os.Getenv("GRUP_CHAT"), 10, 64)
var grupmessageCorona = tgbotapi.NewMessage(grupchatid, "")
var grupmessageWeather = tgbotapi.NewMessage(grupchatid, "")

func cronfunction (c *cron.Cron, conn *redis.Conn, bot *tgbotapi.BotAPI){
	log.Info("Cron Job Started . . . ")
	c.AddFunc("*/1 * * * *", func() { redisHealthCheck(conn)})
	c.AddFunc("*/15 17-18 * * *", func() { coronaUpdate(bot) })
	c.AddFunc("0 6 * * *", func() { weatherUpdate(bot) })
	c.Start()
}

func redisHealthCheck(conn *redis.Conn){
	log.Info("Healthcheck Status: "+ rdc.RedisClientPing(*conn))
}

func coronaUpdate(bot *tgbotapi.BotAPI){
	getHeaderMessage(&grupmessageCorona)
	grupmessageCorona.Text=grupmessageCorona.Text+corona("/corona indonesia")
	grupmessageCorona.ParseMode="markdown"
	_,err:=bot.Send(grupmessageCorona)
	if err!=nil{
		log.Error(err)
	}else{
		log.Info("Message Status: Delivered")
	}
}

func weatherUpdate(bot *tgbotapi.BotAPI){
	getHeaderMessage(&grupmessageWeather)
	grupmessageWeather.Text=grupmessageCorona.Text+cuaca("/cuaca Pasuruan")
	_,err:=bot.Send(grupmessageCorona)
	if err!=nil{
		log.Error(err)
	}else{
		log.Info("Message Status: Delivered")
	}
}

func getHeaderMessage(grupmessage *tgbotapi.MessageConfig){
	h,m,_:=time.Now().Clock()
	jam:=strconv.Itoa(h)
	menit:=strconv.Itoa(m)
	if m==0{
		menit=menit+"0"
	}else if m<10{
		menit="0"+menit
	}
	grupmessage.Text="*Pukul " + jam +":"+ menit +"*\n"
	grupmessage.ParseMode="markdown"
}