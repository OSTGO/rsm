package operate

import (
	"fmt"
	"net/http"
	"strings"
)

func ServerJiangPush(msg string, sendKeys ...string) {
	msg = choreMsg(msg)
	for _, sendKey := range sendKeys {
		url := `https://sc.ftqq.com/` + sendKey + `.send?title=MetaRssMsg&desp=` + msg
		_, err := http.Get(url)
		if err != nil {
			fmt.Println("push error:", err)
		}
	}
}

func choreMsg(msg string) string {
	return strings.Replace(msg, "\n", "%0D%0A", -1)
}
