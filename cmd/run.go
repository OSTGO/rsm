package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"rsm/operate"
)

var jiangHooks []string
var manbuHooks []string
var telegramHooks []string
var dingTalkHooks []string
var serverPort uint16
var gapTime int64
var defaultGapTime int64 = 24
var pages []string
var defaultPage = []string{"rss.html"}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "ru", "R", "RU", "RUN"},
	Short:   "run rsm",
	Long: "`-c,--cfg [rssConfig list]` defaults to `rss.json`, you can specify the configuration file, the configuration file defaults to `rss.json`, you can customize the configuration name, supports the following formats `.toml` ` .yaml` `.toml` `hcl` etc." +
		"`-p,--page [generate html address list]` defaults to rss.html, you can specify the file name to copy" +
		"`-s,--server [port]` can specify the UI interface for starting the server, online preview and addition, deletion, modification and checking." +
		"`-d,--ding [dingTalkHookPath list]` Dingding sending address" +
		"`-t,--telegram [telegramHookPath list]` can specify the telegram receiving address" +
		"`-j,--jiang [serverJiangHookPath list]` can specify the server sauce address, and then push WeChat and other places" +
		"`-m,--manbu [manbuHookPath list]` can specify the receiving address of the `Walking Technology Life` public account, follow the public account to receive messages" +
		"`-g,--gapTime [time interval to get, default 24 hours]` specifies how many hours before the current start, todo: any time range",
	Run: func(cmd *cobra.Command, args []string) {
		op := operate.RssCfgMap{}
		op.MergeCfgFile(rssCfgFiles...)
		RawMsg := operate.GetPostInfo(op, gapTime)
		markDownMsg := operate.SetMarkdown(RawMsg)
		if len(jiangHooks) > 0 {
			operate.ServerJiangPush(markDownMsg, jiangHooks...)
		}
		if len(dingTalkHooks) > 0 {
			operate.DingTalkPush(markDownMsg, dingTalkHooks...)
		}
		if len(manbuHooks) > 0 {
			fmt.Println("wrong token")
		}
		if len(telegramHooks) > 0 {
			operate.TelegramPush(markDownMsg, telegramHooks...)
		}
		if len(pages) > 0 {
			operate.WriteToFile(operate.SetUrl(RawMsg), pages...)
		}
		if serverPort != 0 {
			_ = operate.Server(serverPort, operate.SetUrl(RawMsg))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringSliceVarP(&jiangHooks, "jiang", "j", []string{}, "")
	runCmd.Flags().StringSliceVarP(&manbuHooks, "manbu", "m", []string{}, "")
	runCmd.Flags().StringSliceVarP(&telegramHooks, "telegram", "t", []string{}, "")
	runCmd.Flags().StringSliceVarP(&dingTalkHooks, "dingtalk", "d", []string{}, "")
	runCmd.Flags().Int64VarP(&gapTime, "gaptime", "g", defaultGapTime, "")
	runCmd.Flags().Uint16VarP(&serverPort, "server", "s", 0, "")
	runCmd.Flags().StringSliceVarP(&pages, "page", "p", defaultPage, "")
}
