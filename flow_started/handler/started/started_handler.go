package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/content"
	"tel-mo/flow_started/schema/hello"
)

func StartHandle(c tele.Context) error {
	keyboard := hello.InitSchemaHelloBtn()
	return c.Send(content.FirstMessage, keyboard)
}
