package operate

import (
	"fmt"
	dingTalkBot "github.com/JetBlink/dingtalk-notify-go-sdk"
	"strings"
)

var colonSeparator = ":"

func tokenSplit(str string) (id, secret string) {
	strList := strings.Split(str, colonSeparator)
	id, secret = strList[0], strList[1]
	return
}

func DingTalkPush(msg string, tokens ...string) {
	for _, token := range tokens {
		robot := dingTalkBot.NewRobot(tokenSplit(token))
		err := robot.SendMarkdownMessage("MetaRssMsg", msg, []string{}, false)
		if err != nil {
			fmt.Println("DingTalk push error because: ", err)
		}
	}
}
