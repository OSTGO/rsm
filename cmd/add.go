package cmd

import (
	"rsm/operate"

	"github.com/spf13/cobra"
)

var recordName string
var recordUrl string

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"A", "ADD", "Add", "a"},
	Short:   "Add a record to the configuration file!",
	Long: `Add a record to one or multi configuration files. For example:
			multi cfg file cmd: rsm add -n LT'sBlog -u https://longtaofsd.fun/atom.xml -c ./config1.json,/etc/config2.yaml,~/temp/config3.toml
			one cfg file cmd: rsm add -n LT'sBlog -u https://longtaofsd.fun/atom.xml -c config.json
			[name] and [url] are required parameters, [cfg]are not required parameters, default rss cfg file is ./rss.yaml
			`,
	Run: func(cmd *cobra.Command, args []string) {
		record := map[string]string{recordName: recordUrl}
		operate.FileAddRecords(record, rssCfgFiles...)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&recordName, "name", "n", "", "")
	addCmd.Flags().StringVarP(&recordUrl, "url", "u", "", "")
	err := addCmd.MarkFlagRequired("url")
	if err != nil {
		return
	}
	err = addCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
}
