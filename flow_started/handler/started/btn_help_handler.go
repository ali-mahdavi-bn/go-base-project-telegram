package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/schema/started"
)

func BtnHelpHandler(c tele.Context) error {
	keyboard := started.InitSchemaHelloSelector()
	return c.Send("kooomakkkkk", keyboard)
}
