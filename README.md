# Bot Telegram

Feature: 
- Key / Value Store
- Check weather current condition from openweatherapi

Those 2 feature consume my 2 module 
- Redis Client from my github : https://github.com/guaychou/redisClient
- Openweatherapi client to create GET request in openweatherapi : https://github.com/guaychou/openweatherapi

### How to Use

Before you run this client example you must have three component
1. Bot telegram token (you can get it from BotFather)
2. Redis URI (example : redis://username:password@hostname:port)
3. Openweather API Token (you can get it from openweatherapi.org)

After that you must you ___environment variable___ to get it work

Default Key | Description |
--- | --- |
BOTCHOU_TOKEN_API | Bot telegram token
OWM_TOKEN_API | Open Weather token api
REDIS_URL | Redis URL
USERNAME_TELEGRAM | Your username telegram (This variable will control who can flush your redis server)
You can modify the key as you wish.

In my case, i deploy my bot into heroku platform with free usage. Some addons like ***redis*** will sleep after 30 minute, and if you try to set some value while the redis is sleeping, the application will crashed, ***Restarted***, so i create a cron job, to ping redis server every minute.  