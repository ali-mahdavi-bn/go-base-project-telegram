package started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/schema/hello"
)

func BtnHelpHandler(c tele.Context) error {
	keyboard := hello.InitSchemaHelloSelector()
	return c.Send("kooomakkkkk", keyboard)
}
