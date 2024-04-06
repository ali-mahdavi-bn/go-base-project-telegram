package keyboard

import (
	tele "gopkg.in/telebot.v3"
	"tel-mo/flow_started/content"
)

var (
	Selector  = &tele.ReplyMarkup{}
	Selector2 = &tele.ReplyMarkup{}

	SelectorPrev = Selector.Data(content.SelectorPrev, content.SelectorPrevText)
	SelectorNext = Selector2.Data(content.SelectorNext, content.SelectorNextText)
)
