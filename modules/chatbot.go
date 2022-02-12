package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/amarnathcjd/yoko/modules/db"
	tb "gopkg.in/telebot.v3"
)

var (
	MediaRe = regexp.MustCompile(`<image>(.*?)</image>`)
	BtnRe   = regexp.MustCompile(`<button>(.*?)</button>`)
	UrlRe   = regexp.MustCompile(`<url>(.*?)</url>`)
	TextRe  = regexp.MustCompile(`<text>(.*?)</text>`)
)

func Chat_bot(c tb.Context) error {
	is_chat := false
	if c.Message().IsReply() && c.Message().ReplyTo.Sender.ID == BOT_ID {
		is_chat = true
	} else if strings.Contains(strings.ToLower(c.Message().Text), "mika") {
		is_chat = true
	}
	if !is_chat {
		return nil
	} else if strings.HasPrefix(c.Message().Text, "/") || strings.HasPrefix(c.Message().Text, "!") {
		return nil
	}
	text := strings.ReplaceAll(c.Message().Text, "mika", "kuki")
	url_q := "https://icap.iconiq.ai/talk?&botkey=icH-VVd4uNBhjUid30-xM9QhnvAaVS3wVKA3L8w2mmspQ-hoUB3ZK153sEG3MX-Z8bKchASVLAo~&channel=7&sessionid=482070240&client_name=uuiprod-un18e6d73c-user-19422&id=true"
	req, err := http.PostForm(url_q, url.Values{"input": {text}})
	if err != nil {
		c.Reply("Chatbot API is down!")
	}
	defer req.Body.Close()
	var resp mapType
	json.NewDecoder(req.Body).Decode(&resp)
	msg := resp["responses"].([]interface{})[0].(string)
	fmt.Println(msg)
	ExtractMeta(msg)
	pattern := regexp.MustCompile(`<image>.+</image>`)
	media := pattern.FindAllStringSubmatch(msg, -1)
	yt := regexp.MustCompile(`<card>.+</card>`).FindAllStringSubmatch(msg, -1)
	btn := regexp.MustCompile(`<button>.+</button>`).FindAllStringSubmatch(msg, -1)
	if yt != nil {
		Parse_ai_msg(c, msg, "youtube")
		return nil
	} else if btn != nil {
		Parse_ai_msg(c, msg, "button")
		return nil
	}
	if media != nil {
		if len(media) != 0 {
			file := strings.ReplaceAll(strings.ReplaceAll(media[0][0], "<image>", ""), "</image>", "")
			if strings.Contains(file, "pandorabots") {
				f := strings.SplitN(media[0][0], "<image>", -1)
				fl := f[len(f)-1]
				file = strings.ReplaceAll(strings.ReplaceAll(fl, "</image>", ""), "</image>", "")

			}
			if strings.HasSuffix(file, "jpg") || strings.HasSuffix(file, "png") {
				c.Reply(&tb.Photo{File: tb.FromURL(file)})
			} else {
				c.Reply(&tb.Animation{File: tb.FromURL(file)})
			}
		}
	}
	var message string
	for _, m := range media {
		message = strings.Replace(message, m[0], "", -1)
	}

	if strings.Contains(message, "<split>") {
		message = strings.ReplaceAll(strings.ReplaceAll(message, "<split>", ""), "</split>", "")
	}
	message = strings.ReplaceAll(strings.ReplaceAll(message, "kuki", "mika"), "Kuki", "Mika")
	c.Bot().Notify(c.Chat(), tb.Typing)
	fmt.Println(message)
	return c.Reply(message)
}

func Chatbot_mode(c tb.Context) error {
	if c.Message().Private() {
		c.Reply("This command is made for group chats only.")
		return nil
	}
	args := c.Message().Payload
	if args == string("") {
		mode := db.Get_chatbot_mode(c.Chat().ID)
		if mode {
			c.Reply("AI chatbot is currently <b>enabled</b> for this chat.")
			return nil
		} else {
			c.Reply("AI chatbot is currently <b>disabled</b> for this chat.")
			return nil
		}
	}
	if stringInSlice(strings.ToLower(args), []string{"enable", "on", "yes", "y"}) {
		c.Reply("<b>Enabled</b> AI chatbot for this chat.")
		db.Set_chatbot_mode(c.Chat().ID, true)
	} else if stringInSlice(strings.ToLower(args), []string{"disable", "off", "no", "n"}) {
		c.Reply("<b>Disabled</b> AI chatbot for this chat.")
		db.Set_chatbot_mode(c.Chat().ID, false)
	} else {
		c.Reply("Your input was not recognised as one of: yes/no/y/n/on/off")
	}
	return nil
}

func Parse_ai_msg(c tb.Context, t string, mode string) {
	if mode == "youtube" {
		title := regexp.MustCompile(`<title>.+</title>`).FindAllStringSubmatch(t, -1)
		subtitle := regexp.MustCompile(`<subtitle>.+</subtitle>`).FindAllStringSubmatch(t, -1)
		msg := ""
		if subtitle != nil {
			msg = strings.ReplaceAll(strings.ReplaceAll(subtitle[0][0], "<subtitle>", ""), "</subtitle>", "")
		} else if title != nil {
			msg = strings.ReplaceAll(strings.ReplaceAll(title[0][0], "<title>", ""), "</title>", "")
		}
		buttons := regexp.MustCompile(`<button>.+</button>`).FindAllStringSubmatch(t, -1)
		if buttons != nil {
			txt := regexp.MustCompile(`<text>.+</text>`).FindAllStringSubmatch(t, -1)
			text := strings.ReplaceAll(strings.ReplaceAll(txt[0][0], "<text>", ""), "</text>", "")
			btn_url := regexp.MustCompile(`<url>.+</url>`).FindAllStringSubmatch(t, -1)
			url := strings.ReplaceAll(strings.ReplaceAll(btn_url[0][0], "<url>", ""), "</url>", "")
			menu.Inline(menu.Row(menu.URL(text, url)))
		}
		fl := regexp.MustCompile(`<image>.+</image>`).FindAllStringSubmatch(t, -1)
		if fl != nil {
			file := strings.ReplaceAll(strings.ReplaceAll(fl[0][0], "<image>", ""), "</image>", "")
			c.Reply(&tb.Photo{File: tb.FromURL(file), Caption: msg}, menu)
		} else {
			c.Reply(msg, menu)
		}
		final_msg := strings.SplitN(t, "</card>", -1)
		final_msg_to_send := final_msg[len(final_msg)-1]
		final_msg_to_send = strings.ReplaceAll(strings.ReplaceAll(final_msg_to_send, "kuki", "yoko"), "Kuki", "Yoko")
		c.Reply(final_msg_to_send)
	} else if mode == "button" {
		text := strings.SplitN(t, "<", 2)[0]
		btn_txt := regexp.MustCompile(`<text>.+</text>`).FindAllStringSubmatch(t, -1)
		button_text := "Link"
		if btn_txt != nil {
			button_text = strings.ReplaceAll(strings.ReplaceAll(btn_txt[0][0], "<text>", ""), "</text>", "")
		}
		btn_url := regexp.MustCompile(`<url>.+</url>`).FindAllStringSubmatch(t, -1)
		url := strings.ReplaceAll(strings.ReplaceAll(btn_url[0][0], "<url>", ""), "</url>", "")
		menu.Inline(menu.Row(menu.URL(button_text, url)))
		c.Reply(text, menu)
	}

}

func ExtractMeta(t string) {
	var Med []string
	for _, x := range MediaRe.FindAllStringSubmatch(t, -1) {
		Med = append(Med, x[1])
		t = strings.Replace(t, x[0], "", 1)
	}
	var Btn map[string]string
	for _, y := range BtnRe.FindAllStringSubmatch(t, -1) {
		url, text := "", "Link"
		fmt.Println(y)
		for _, x := range UrlRe.FindAllStringSubmatch(y[1], -1) {
			url = x[1]
		}
		for _, x := range TextRe.FindAllStringSubmatch(y[1], -1) {
			text = x[1]
		}
		Btn[url] = text
	}
	fmt.Println(Btn, Med, t)
}
