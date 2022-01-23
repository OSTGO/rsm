[🇨🇳](readme/README_CN.md)[中文](readme/README_CN.md)

# rsm :sparkles:

RSS master[rsm] is a RSS subscription function aggregation tool, You can use it easily!

# How to start？ :helicopter:

Start rsm with `rsm run`

- `-c,--cfg [rssConfig list]` defaults to `rss.json`, you can specify the configuration file, the configuration file
  defaults to `rss.json`, you can customize the configuration name, supports the following
  formats `.toml` ` .yaml` `.toml` `hcl` etc.
- `-p,--page [generate html address list]` defaults to rss.html, you can specify the file name to copy
- `-s,--server [port]` can specify the UI interface for starting the server, online preview and addition, deletion,
  modification and checking.
- `-d,--ding [dingTalkHookPath list]` Dingding sending address
- `-t,--telegram [telegramHookPath list]` can specify the telegram receiving address
- `-j,--jiang [serverJiangHookPath list]` can specify the server sauce address, and then push WeChat and other places
- `-m,--manbu [manbuHookPath list]` can specify the receiving address of the `Walking Technology Life` public account,
  follow the public account to receive messages
- `-g,--gapTime [time interval to get, default 24 hours]` specifies how many hours before the current start, todo: any
  time range

# How to operate rss List :fuelpump:

You can modify the `rss.json` file directly, or you can customize the json subscription list and
use `-c,--cfg [rssConfig]` to specify the subscription list you want to use at startup

You can also manipulate the subscription list with the following commands

- `rsm list -c [rssConfig]` displays a list of subscriptions
- `rsm add -n [name] -u [url] -c [rssConfig]` add subscription
- `rsm delete -n [name list] -c [rssConfig]` delete subscription
- `rsm merge -c [merged rssConfig list] -o [merged rssConfig]` merge subscription list files

# Install :package:

Download the version that works for u

- [Mac-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-mac-amd64)
- [Windows-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-win-amd64.exe)
- [linux-arm](https://github.com/metaRSS/rsm/releases/download/v/rsm-linux-arm64)
- [linux-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-linux-amd64)

# Deploy :rocket:

## local deployment

Download the executable file to the deployment server and set it to run regularly

## github action deployment

Clone this repo, modify the subscription list, and modify the startup parameters in daily.yaml

DingTalk example:
![dingTalkSample](https://gcore.jsdelivr.net/gh/Longtao-W/pics@main/rsm/dingTalkSample.jpg)

Example of server sauce:
![serverJiangSample](https://gcore.jsdelivr.net/gh/Longtao-W/pics@main/rsm/serverJiangSample.jpg)

# license

Use GPL license

# Todo list :clipboard:

[]UI interaction to increase rss
[]Subscribe rss by email
[]Rss to mail
