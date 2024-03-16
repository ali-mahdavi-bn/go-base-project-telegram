package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/schema/hello/keyboard"
)

func SelectorNextHandler(c tele.Context) error {
	return c.Edit("hear edite help", keyboard.Selector)
}
