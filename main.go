
package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	rdc "github.com/guaychou/redisClient"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)



func main() {
	// Create redis connection
	pool:= rdc.NewPool(20,20)
	conn := pool.Get()
	defer conn.Close()


	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOTCHOU_TOKEN_API"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	// Create cronjob
	c := cron.New()
	cronfunction(c, &conn, bot)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text="I'm a babu of lordchou thanks for add me,\nlordchou LinkedIn: https://www.linkedin.com/in/kevin-jonathan-harnanta-b0745216b/ \nSend /help for more information."
			case "help":
				msg.Text = "Type /wifirumah : for Password Wifi-Rumah \nor\nType /wifikos : for Password Wifi Kos\nType /cuaca : to get the latest weather conditions"
			case "sayhi":
				msg.Text = "Hi "+ update.Message.From.UserName + ", Nice to meet you"
			case "status":
				msg.Text = "I'm ok."
			case "wifirumah":
				msg.Text="Password: "+ rdc.RedisClientGet(conn,"wifirumah")
			case "wifikos":
				msg.Text="Password: "+ rdc.RedisClientGet(conn,"wifikos")
			case "set":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=3{
					msg.Text="Some argument is missing. \nUse /set <key> <value> to store data."
				}else if len(split)==3{
					key:=split[1]
					value:=split[2]
					msg.Text=rdc.RedisClientSet(conn,key,value)
				}
			case "get":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=2{
					msg.Text="Some argument is missing. \nUse /get <key> to get the value."
				}else if len(split)==2 {
					key:=split[1]
					msg.Text=rdc.RedisClientGet(conn,key)
				}
			case "flush":
				if update.Message.From.UserName==os.Getenv("USERNAME_TELEGRAM"){
					msg.Text=rdc.RedisClientFlush(conn)
				}else{
					msg.Text="Forbidden Status: You are not my lord !!!"
				}
			case "del":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=2{
					msg.Text="Some argument is missing. \nUse /get <key> to get the value."
				}else if len(split)==2 {
					key:=split[1]
					msg.Text=rdc.RedisClientDelete(conn,key)
				}
			case "cuaca":
				msg.Text=cuaca(update.Message.Text)
			case "corona":
				msg.Text=corona(update.Message.Text)
			default:
				msg.Text = "I don't know that command"
			}
			msg.ParseMode="markdown"
			bot.Send(msg)
		}

	}
}