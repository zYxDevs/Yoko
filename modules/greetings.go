package modules

import (
	"fmt"
	"strconv"

	"github.com/amarnathcjd/yoko/modules/db"
	tb "gopkg.in/tucnak/telebot.v3"
)

func Welcome_set(c tb.Context) error {
	if c.Message().Private() {
		c.Reply("This command is made to be used in group chats.")
		return nil
	}
	if c.Message().Payload == string("") {
		text, file, mode := db.Get_welcome(c.Chat().ID)
		c.Reply(fmt.Sprintf("<b>Greetings config in this chat</b>:\n- Should greet new members: <code>%s<code>\n- Delete old welcome message: <code>%s</code>\n- Delete welcome service: <code>%s</code>\n\nWelcome message:", strconv.FormatBool(mode), "True", "True"))
		if mode {
			unparse_message(file, text, c.Message())
		}
	}
	return nil
}
