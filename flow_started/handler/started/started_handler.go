package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/content"
	"tel-mo/flow_started/schema/started"
)

func StartHandle(c tele.Context) error {
	keyboard := started.InitSchemaHelloBtn()
	return c.Send(content.FirstMessage, keyboard)
}
