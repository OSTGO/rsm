package operate

import (
	"fmt"
	"net/http"
)

func TelegramPush(msg string, sendKeys ...string) {
	msg = choreMsg(msg)
	for _, sendKey := range sendKeys {
		token, chatId := tokenSplit(sendKey)
		url := `https://api.telegram.org/bot` + token + `/sendMessage?chat_id=` + chatId + `&text=` + msg
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("push error:", err)
		}
	}
}
