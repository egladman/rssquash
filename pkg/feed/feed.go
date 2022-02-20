package feed

import (
	"bufio"
	"bytes"
	"fmt"
	"path"
	"os"
	"sort"
	"text/template"

	"github.com/mmcdole/gofeed"
	"github.com/egladman/rssquash/internal/utils"
	"github.com/egladman/rssquash/pkg/defaults"
)

var (
	FeedItems []*gofeed.Item
)

var TemplateFuncMap = template.FuncMap{
	"GetCurrentTime": utils.GetCurrentTime,
	"GetFeedBaseUrl": utils.GetFeedBaseUrl,
	"GetFeedPrefixUrl": utils.GetFeedPrefixUrl,
	"GetFeedBaseName": utils.GetFeedBaseName,
	"GetFeedTitle": utils.GetFeedTitle,
}

func Fetch(s string) (*gofeed.Feed, error){
	fp := gofeed.NewParser()

	fd, err := fp.ParseURL(s)
	if err != nil {
		return fd, err
	}

	return fd, nil
}

func Render(items []*gofeed.Item) (bytes.Buffer, error) {
	// The basename of the value passed to ParseFiles must match the template
	// name. Otherwise an error will get thrown. This because New gets called
	// once more underhood
	n := path.Base(defaults.FeedTemplatePath)
	t, err := template.New(n).Funcs(TemplateFuncMap).ParseFiles(defaults.FeedTemplatePath)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("Failed to parse template: %v", err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, items)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("Failed to execute template: %v", err)
	}

	return buf, nil
}

func Read(p string) ([]*gofeed.Item, error) {
	file, err := os.Open(p)
	if err != nil {
		return []*gofeed.Item{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fd, err := Fetch(s)
		if err != nil {
			// TODO Add bool to optionally error out
			fmt.Printf("Failed to fetch feed '%s'. Skipping\n", s)
			continue
		}
		FeedItems = append(FeedItems, fd.Items...)
	}

	if err := scanner.Err(); err != nil {
		return []*gofeed.Item{}, fmt.Errorf("Failed to read file: %v", err)
	}

	sort.SliceStable(FeedItems, func(i, j int) bool {
	 	ti := FeedItems[i].PublishedParsed
	 	tj := FeedItems[j].PublishedParsed
	 	return ti.Before(*tj)
	})

	return FeedItems, nil
}

func Generate(s string) (bytes.Buffer, error) {
	items, err := Read(s)
	if err != nil {
		return bytes.Buffer{}, err
	}

	buf, err := Render(items)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return buf, nil
}
