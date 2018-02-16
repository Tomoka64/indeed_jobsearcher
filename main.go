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
	Jobbs []Job
}

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

		goQuery(url, number)
	}
	app.Run(os.Args)
}

func (j *Jobs) Add(item Job) []Job {
	j.Jobbs = append(j.Jobbs, item)
	return j.Jobbs
}

func goQuery(url string, num int) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	counter := num
	count := 0
	for count < counter {
		doc.Find("tr").Each(func(i int, s *goquery.Selection) {

			title := s.Find("a").Text()
			location := s.Find("td").Text()

			location = strings.Join(strings.Fields(location), " ")

			jobby := &Job{
				Title:   title,
				Summary: location,
			}
			// items := []Job{}
			// Job := Jobs{items}
			// a := Job.Add(*jobby)
			// fmt.Printf("%v\n", a)
			color.Red("Job title %d: %s ", i, jobby.Title)
			color.Yellow("Detail: %s \n", jobby.Summary)

			count++
		})
	}

}

func buildURL(key1 string) string {
	return fmt.Sprintf("http://www.expatengineer.net/?query=%s", url.QueryEscape(key1))
}
