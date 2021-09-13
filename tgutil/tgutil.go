package tgutil

import (
	"bytes"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xiaoyueya/chain-sdk/bwlog"
	"math/big"
)

// SendTgMsg /** 发送tg消息，如果消息中包含链接，不展示预览，省得占屏幕；每个消息可以加一个标题，当然你也可以把消息头当消息 **/
func SendTgMsg(botKey string, chatId string, title string, msgLines ...string) {
	go func() {
		chatIdInt64, _ := big.NewInt(0).SetString(chatId, 0)
		bot, err := tgbotapi.NewBotAPI(botKey)
		if err != nil {
			bwlog.Logger.Errorf("new tg bog error=%v\n", err)
			return
		}

		bufBuf := new(bytes.Buffer)
		bufBuf.WriteString(fmt.Sprintf("<pre><b>%s</b></pre>\n", title))
		bufBuf.WriteString("<pre><b>============================</b></pre>\n")
		for _, line := range msgLines {
			bufBuf.WriteString(fmt.Sprintf("<pre>%s</pre>\n", line))
		}
		plainText := bufBuf.String()

		msg := tgbotapi.NewMessage(chatIdInt64.Int64(), plainText)

		//msg.ParseMode = "markdown"
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true
		reply, err := bot.Send(msg)
		if err != nil {
			bwlog.Logger.Errorf("tg bog send msg error=%v\n", err)
			return
		}
		bwlog.Logger.Debugf("reply=%s\n", reply.Text)
	}()
}
