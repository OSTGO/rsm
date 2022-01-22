package cmd

import (
	"github.com/spf13/cobra"
	"rsm/operate"
)

var delNames []string
var delUrls []string

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "D", "Delete", "del", "Del"},
	Short:   "Delete a record to the configuration file!",
	Long: `Delete a record to one or multi configuration files. For example:
			multi cfg file cmd with record name: rsm del -n name1,name2 -u https://longtaofsd.fun/atom.xml -c ./config1.json,/etc/config2.yaml,~/temp/config3.toml
			one cfg file cmd with record name: rsm del -n name1 -u https://longtaofsd.fun/atom.xml -c config.json
			multi cfg file cmd with record url: rsm del -n name1,name2,name3,name4,name5,name6,namex -u https://longtaofsd.fun/atom.xml -c ./config1.json,/etc/config2.yaml,~/temp/config3.toml
			At least one of [name] or [ur]l is required, [cfg]are not required parameters, default rss cfg file is ./rss.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		operate.FileDelRecords(delNames, delUrls, rssCfgFiles...)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringSliceVarP(&delNames, "name", "n", []string{}, "")
	deleteCmd.Flags().StringSliceVarP(&delUrls, "url", "u", []string{}, "")
}
