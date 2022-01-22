package operate

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"time"
)

const htmlFront = "<!DOCTYPE html>\n<html lang=\"en\">\n<style type=\"text/css\">\n    .ime {\n        width: 100%;\n        height: 50px;\n    }\n\n    .ime:hover {\n        width: 100%;\n        height: 600px;\n    }\n</style>\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n</head>\n<body style=\"border: 1px solid #000000; width: 90%; height: 100%;margin: 0 auto;\">\n"
const htmlBack = "\n</body>"

func getTime(gapTime time.Duration) (int64, int64) {
	now := time.Now().UTC()

	startTime := now.Add(-gapTime * time.Hour)
	start := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 0, 0, 0, now.Location()).Unix()
	end := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).Unix()
	return start, end
}

type outComeRecord struct {
	title string
	link  string
}

func GetPostInfo(rss RssCfgMap, gapTime int64) []outComeRecord {
	var msg = make([]outComeRecord, 0)
	startTime, endTime := getTime(time.Duration(gapTime))
	fp := gofeed.NewParser()
	for name, url := range rss {
		msg = append(msg, outComeRecord{name, ""})
		feed, err := fp.ParseURL(url)
		if err != nil {
			fmt.Print(err.Error())
		} else {
			for _, item := range feed.Items {
				if item.PublishedParsed != nil && item.PublishedParsed.Unix() >= startTime && item.PublishedParsed.Unix() < endTime {
					msgItem := outComeRecord{title: item.Title, link: item.Link}
					msg = append(msg, msgItem)
				}
			}
		}
	}
	return msg
}

func SetMarkdown(es []outComeRecord) string {
	var a string
	for _, e := range es {
		a = fmt.Sprint(a, fmt.Sprintf("[%s](%s)", e.title, e.link), "\n\n")
	}
	return a
}

func SetUrl(es []outComeRecord) string {
	var a string
	for _, e := range es {
		a = fmt.Sprint(a, fmt.Sprintf("<a href=\"%s\">%s</a><br><iframe  class=\"ime\" src=\"%s\"  allowfullscreen=\"true\">  </iframe><br>", e.link, e.title, e.link), "\n")
	}
	return htmlFront + a + htmlBack
}
