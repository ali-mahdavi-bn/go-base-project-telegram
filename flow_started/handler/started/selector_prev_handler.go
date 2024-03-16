package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/schema/hello"
)

func SelectorPrevHandler(c tele.Context) error {
	keyboard := hello.InitSchemaBtnPrev()
	return c.Edit("hear edite help ascdascasdc", keyboard)
}
