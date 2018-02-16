package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type Job struct {
	Title   string
	Summary string
}

type Jobs struct {
	Jobs []Job
}

// // type QueryResult struct {
// // 	Title string
// // 	Job   []*job
// // }
//
func main() {
	app := cli.NewApp()
	app.Name = "indeed"
	app.Usage = "Command Line for scraping information of available jobs"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "number, n",
			Value: 10,
			Usage: "number of output lines",
		},
		cli.StringFlag{
			Name:  "word, w",
			Value: "engineer",
			Usage: "searching word",
		},
	}
	app.Action = func(c *cli.Context) {
		number := c.Int("number")
		keyword := c.String("word")
		url := buildURL(keyword)
		// jobs := []Job{}

		goQuery(url, number)
	}
	app.Run(os.Args)
}

// (QueryResult, error)

// func main() {
// 	var keyword = flag.String("key", "engineer", "keyword")
// 	url := buildURL(*keyword)
// 	goQuery(url, 10)
// }
func goQuery(url string, num int) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	// counter, _ := strconv.Atoi(num)
	// count := 0
	// for count < counter {
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("a").Text()
		location := s.Find("td").Text()

		location = strings.Join(strings.Fields(location), " ")

		jobby := &Job{
			Title:   title,
			Summary: location,
		}
		color.Red("Job title %d: %s ", i, jobby.Title)
		color.Yellow("Detail: %s \n", jobby.Summary)

		// title := s.Find("i").Text()
		// fmt.Printf("Review %d: %s - %s\n", i, title)
		// count++
	})
	// }

}

//
// func (j *Jobs) Graph() {
// 	for i := range j.Jobs {
// 		color.Red("Job title %d: %s ", i.Title, title)
// 		color.Yellow("Job Location %d: %s \n", i, location)
// 	}
// }

func buildURL(key1 string) string {
	return fmt.Sprintf("http://www.expatengineer.net/?query=%s", url.QueryEscape(key1))

}
