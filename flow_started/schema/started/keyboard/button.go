package keyboard

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/content"
)

var (
	Menu        = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnHelp     = Menu.Text(content.BtnHelp)
	BtnSettings = Menu.Text(content.BtnSettings)
)
