package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/egladman/rssquash/pkg/feed"
)

var (
	FeedSourcePath   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rssquash",
	Short: "RSS/Atom/JSON feed aggregator",
	Long: `Consolidate multiple RSS/Atom/JSON feeds into one easily consumable Atom feed`,

	Run: func(cmd *cobra.Command, args []string) {
		buf, err := feed.Generate(FeedSourcePath)
		if err != nil {
			panic(err)
		}
		fmt.Println(buf.String())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&FeedSourcePath, "source", "s", "", "Path to new line delimited list of feed URLs")
	rootCmd.MarkFlagRequired("source")
}
