package application

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started"
)

func LoadHandler(bot *tele.Bot) {
	for command, handler := range flow_started.LoadHandlerHello() {
		bot.Handle(command, handler)
	}
}
