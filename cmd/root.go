/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"path"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/mmcdole/gofeed"
	"github.com/egladman/rssquash/internal/utils"
)

const (
	FeedTemplatePathDefault = "configs/feed.atom.tmpl"
	FeedSourcePathDefault   = "examples/feeds.list"
)

var (
	FeedTemplatePath string
	FeedSourcePath   string
	FeedItems        []*gofeed.Item
)

var TemplateFuncMap = template.FuncMap{
	"getCurrentTime": utils.GetCurrentTime,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rssquash",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		readFeedFile(FeedSourcePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rssquash.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&FeedSourcePath, "source", "s", FeedSourcePathDefault, "Source file to read from")
	rootCmd.PersistentFlags().StringVarP(&FeedTemplatePath, "template", "t", FeedTemplatePathDefault, "Source file to read from")
}

func fetchFeed(s string) (*gofeed.Feed, error){
	fp := gofeed.NewParser()

	//fmt.Println("Fetching ", s)
	fd, err := fp.ParseURL(s)
	if err != nil {
		return fd, err
	}

	return fd, nil
}

func consumeFeed(fd *gofeed.Feed) []*gofeed.Item {
	//	for _, item := range fd.Items {
	//	fmt.Println(item)
	//}

	// Todo filtering?

	return fd.Items
}

func renderFeed(it []*gofeed.Item) {
	// The basename of the value passed to ParseFiles must match the template
	// name. Otherwise an error will get thrown. This because New gets called
	// once more underhood
	n := path.Base(FeedTemplatePath)
	t, err := template.New(n).Funcs(TemplateFuncMap).ParseFiles(FeedTemplatePath)

	//t, err := template.ParseFiles(FeedTemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, it)
	if err != nil {
		log.Fatal("Failed to execute template: %v", err)
	}
}

func readFeedFile(p string) {
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		fd, err := fetchFeed(s)
		if err != nil {
			fmt.Printf("Failed to retrieve feed from url '%s'. Skipping", s)
			continue
		}
		items := consumeFeed(fd)
		FeedItems = append(FeedItems, items...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	renderFeed(FeedItems)
}






