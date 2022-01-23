[🇺🇸](../README.md)[English](../README.md)

# rsm :sparkles:

RSS master[rsm], RSS大师 是一个RSS订阅工具，你可以免费且轻松的使用

# 如何启动 :helicopter:

使用`rsm run`启动rsm

- `-c,--cfg [rssConfig列表]`默认为`rss.json`，可指定配置文件,配置文件默认为`rss.json`，可自定义配置名称， 支持以下格式`.toml` `.yaml` `.toml` `hcl`等
- `-p,--page [生成html地址列表]`默认为rss.html,可指定文件名复制
- `-s,--server [port]`可指定启动server可以在线预览以及增删改查什么的UI界面
- `-d,--ding [dingTalkHookPath列表]` 钉钉发送地址
- `-t,--telegram [telegramHookPath列表]` 可指定telegram接收地址
- `-j,--jiang [serverJiangHookPath列表]` 可指定Server酱地址，进而推送的微信等地方
- `-m,--manbu [manbuHookPath列表]` 可指定`漫步科技人生`公众号接收地址，关注公众号可接收消息
- `-g,--gapTime [获取的时间间隔，默认24小时]` 指定从当前开始开始前多少小时的结果，todo：任意时间范围

# 如何增删订阅列表 :fuelpump:

你可以直接修改`rss.json`文件，也可以自定义json订阅列表并在启动时使用`-c,--cfg [rssConfig]`指定你要使用的订阅列表

你也可以使用以下命令操作订阅列表

- `rsm list -c [rssConfig]` 显示订阅列表
- `rsm add -n [name] -u [url] -c [rssConfig]` 增加订阅
- `rsm delete -n [name列表] -c [rssConfig]` 删除订阅
- `rsm merge -c [被合并的rssConfig列表] -o [合并后的的rssConfig]` 合并订阅列表文件

# 安装 :package:

在此下载适用于你的版本

- [Mac-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-mac-amd64)
- [Windows-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-win-amd64.exe)
- [linux-arm](https://github.com/metaRSS/rsm/releases/download/v/rsm-linux-arm64)
- [linux-amd64](https://github.com/metaRSS/rsm/releases/download/v/rsm-linux-amd64)

如果你需要特殊版本，请下载本repo自行编译

# 定时任务部署 :rocket:

## 本地部署

下载可执行文件到部署服务器中,设置定时运行

## github action 部署

克隆本repo,修改订阅列表,修改daily.yaml中的启动参数

钉钉示例:
![dingTalkSample](https://gcore.jsdelivr.net/gh/Longtao-W/pics@main/rsm/dingTalkSample.jpg)

server酱示例:
![serverJiangSample](https://gcore.jsdelivr.net/gh/Longtao-W/pics@main/rsm/serverJiangSample.jpg)

# 许可证

本项目遵循GPL许可，您可以免费的使用和传播，但不可用于商业目的。
