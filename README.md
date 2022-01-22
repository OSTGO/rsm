[🇨🇳](readme/README_CN.md)[中文](readme/README_CN.md)

# rsm :sparkles:

rsm is an RSS subscription function aggregation tool, You can use it easily!

# How to start？ :helicopter:

run command `rsm run`
here is optionals

- `-d,--ding [dingTalkHookPath]` 钉钉发送地址
- `-r,--rss [rssConfig]`默认为`rss.json`，可指定配置文件
- `-h,--html [生成html地址]`默认为out.html,可指定文件名复制
- `-s,--server [port]`可指定启动server可以在线预览以及增删改查什么的UI界面
- `-t,--telegram [telegramHookPath]` 可指定telegram接收地址

# How to operate rss List :fuelpump:

you can directly change `rss.json` or create json file follow the same format; u can also use this command to amend rss
list:
`rsm list` show rss name list
`rsm add -n [name] -p [path]` add rss
`rsm remove [name]` remove rss by name
``
``
``

# Install :package:

Download the version that works for u [Mac-amd64](!mac)|[Windows-amd64](!win)|[linux-arm](!l)|[linux-amd64](!l)
|[linux-x86-32bit](!l)|

# Deploy :rocket:

deploy by local

deploy by github

# license

use GPL license

# Todo list :clipboard:

[]UI interaction to increase rss
[]Subscribe rss by email
[]Rss to mail
[]Push rss to telegram
[]Push rss to WeChat
[]Beautify UI display
