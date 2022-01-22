package cmd

import (
	"github.com/spf13/cobra"
	"rsm/operate"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "LIST", "List", "L"},
	Short:   "Show RSS subscription records",
	Long: `Show RSS subscription records,for example:
rsm list -c ./config1.json,/etc/config2.yaml,~/temp/config3.toml
[cfg]are not required parameters, default rss cfg file is ./rss.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		op := operate.RssCfgMap{}
		op.MergeCfgFile(rssCfgFiles...)
		op.PrintKeyValue(15, 60)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
