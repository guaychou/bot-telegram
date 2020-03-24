# Bot Telegram

Feature: 
- Key / Value Store
- Check weather current condition from openweatherapi
- Check corona status from https://covid19.mathdro.id , and consume client from https://github.com/guaychou/corona-api

Those 2 feature consume my 2 module 
- Redis Client from my github : https://github.com/guaychou/redisClient
- Openweatherapi client to create GET request in openweatherapi : https://github.com/guaychou/openweatherapi

### Requirement

Before you run this client example you must have three component
1. Bot telegram token (you can get it from BotFather)
2. Redis URI (example : redis://username:password@hostname:port)
3. Openweather API Token (you can get it from openweatherapi.org)

After that you must set your ___environment variable___ to get it work

Default Variable | Description |
--- | --- |
BOTCHOU_TOKEN_API | Bot telegram token
OWM_TOKEN_API | Open Weather token api
REDIS_URL | Redis URL
USERNAME_TELEGRAM | Your username telegram (This variable will control who can flush your redis server)

You can modify the variable as you wish.

In my case, i deploy my bot into heroku platform with free usage. Some addons like ***redis*** will sleep after 30 minute, and if you try to set some value while the redis is sleeping, the application will crashed, ***Restarted***, so i create a cron job, to ping redis server every minute.

### How to use : 
/corona countryName : Get current corona virus status from specific country

/cuaca cityName : Get current weather from specific city name  

/set key value : store value to redis

/get key : Get value from specific key