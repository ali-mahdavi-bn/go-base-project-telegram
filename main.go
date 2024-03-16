package main

import (
	"gopkg.in/telebot.v3/middleware"
	"tel-mo/application"
)

func main() {
	bot := application.TelegramBot()

	// middleware
	bot.Use(middleware.Logger())
	bot.Use(middleware.AutoRespond())

	/// Load Handler
	application.LoadHandler(bot)

	// started app
	bot.Start()
}
