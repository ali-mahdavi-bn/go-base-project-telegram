package flow_started

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/handler/started"
	"tel-mo/flow_started/schema/started/keyboard"
)

func LoadHandlerHello() map[any]func(m tele.Context) error {
	commandMap := make(map[any]func(c tele.Context) error)

	// command
	commandMap["/start"] = started.StartHandle

	// button
	commandMap[&keyboard.SelectorPrev] = started.SelectorPrevHandler
	commandMap[&keyboard.SelectorNext] = started.SelectorNextHandler
	commandMap[&keyboard.BtnHelp] = started.BtnHelpHandler
	return commandMap
}
