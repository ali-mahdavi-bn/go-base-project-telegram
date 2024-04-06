package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/schema/started"
)

func SelectorPrevHandler(c tele.Context) error {
	keyboard := started.InitSchemaBtnPrev()
	return c.Edit("hear edite help ascdascasdc", keyboard)
}
