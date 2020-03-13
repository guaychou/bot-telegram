
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gomodule/redigo/redis"
	owm "github.com/guaychou/openweatherapi"
	)

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 20,
		// max number of connections
		MaxActive: 20,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(os.Getenv("REDIS_URL"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func redisClientSet(c redis.Conn, key string, value string) string {
	_, err := c.Do("SET", key, value)
	if err != nil {
		log.Fatal(err)
	}
	return "Value of "+key+" has been stored."
}

func redisClientGet(c redis.Conn,key string) string{
	values, err := redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		return "Key "+key+" does not exist !!! .\nSET first with /set command !!!"
	} else if err != nil {
		log.Fatal(err)
	}
	return values
}
func redisClientFlush(c redis.Conn) string {
	err := c.Send("FLUSHALL")
	if err != nil {
		log.Fatal(err)
	}
	return "All key values has been deleted."
}
func main() {
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()
	api_key:=os.Getenv("OWM_TOKEN_API")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOTCHOU_TOKEN_API"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

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
				msg.Text="Password: "+ redisClientGet(conn,"wifirumah")
			case "wifikos":
				msg.Text="Password: "+ redisClientGet(conn,"wifikos")
			case "set":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=3{
					msg.Text="Some argument is missing. \nUse /set <key> <value> to store data."
				}else if len(split)==3{
					key:=split[1]
					value:=split[2]
					msg.Text=redisClientSet(conn,key,value)
				}
			case "get":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=2{
					msg.Text="Some argument is missing. \nUse /get <key> to get the value."
				}else if len(split)==2 {
					key:=split[1]
					msg.Text=redisClientGet(conn,key)
				}
			case "flush":
				if update.Message.From.UserName==os.Getenv("USERNAME_TELEGRAM"){
					msg.Text=redisClientFlush(conn)
				}else{
					msg.Text="Forbidden Status: You are not my lord !!!"
				}
			case "cuaca":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=2{
					msg.Text="Some argument is missing. \nUse /cuaca <namaKota> to get the value."
				}else if len(split)==2 {
					kota:=split[1]
					result:=owm.GetWeather(kota,api_key)
					if result.Cod!=200{
						msg.Text="City not found."
					}else {
						city:=result.Name
						humidity:=strconv.Itoa(result.Humidity)
						description:=result.Weather[0].Description
						suhu:=fmt.Sprintf("%.2f",result.Temp)
						suhuMaks:=fmt.Sprintf("%.2f",result.Temp_max)
						suhuMin:=fmt.Sprintf("%.2f",result.Temp_min)
						kelembaban:=result.Kelembapan
						msg.Text="Kota: "+city+"\nCuaca: "+description+"\nSuhu: "+suhu+" °C\nSuhu Minimal: "+suhuMin+" °C\nSuhu Maksimal: "+suhuMaks+" °C\nAngka Kelembaban: "+humidity+"\nStatus Kelembaban: "+kelembaban
					}
				}

			default:
				msg.Text = "I don't know that command"
			}
			msg.ParseMode="markdown"
			bot.Send(msg)
		}

	}
}
