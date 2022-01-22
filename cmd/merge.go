package cmd

import (
	"github.com/spf13/cobra"
	"rsm/operate"
)

var rssCfgOut []string

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Aliases: []string{"m", "M", "Merge", "mer", "MER", "Mer"},
	Short:   "Merge and deduplicate record files",
	Long: `Merge and deduplicate record files, for example:
rsm merge -c ./config1.json,/etc/config2.yaml,~/temp/config3.toml -o ./out.yaml,/etc/configOut.json
[cfg] and [out] are not required parameters, default rss cfg file is ./rss.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		op := operate.RssCfgMap{}
		op.MergeCfgFile(rssCfgFiles...)
		op.WriteCfgFileList(rssCfgOut...)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringSliceVarP(&rssCfgOut, "out", "o", defaultRssCfgFiles, "")

}
