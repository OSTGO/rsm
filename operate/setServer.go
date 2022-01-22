package operate

import (
	"fmt"
	"net/http"
	"strconv"
)

type rssIndexStruct struct {
	msg string
}

func (rss *rssIndexStruct) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, rss.msg)
	if err != nil {
		return
	}
}

func Server(port uint16, msg string) bool {
	http.Handle("/", &rssIndexStruct{msg: msg})
	fmt.Println("http://127.0.0.1:" + strconv.Itoa(int(port)))
	err := http.ListenAndServe(":"+strconv.Itoa(int(port)), nil)
	if err != nil {
		fmt.Println("Server failed! because: ", err)
	}
	return true
}
