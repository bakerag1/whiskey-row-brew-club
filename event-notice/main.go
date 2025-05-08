package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	name := os.Args[1]

	if name == "migrate" {
		makeNotices()
		return
	}
	makeEvent(name)

}

func makeEvent(name string) {
	log.Println(name)
	var mtg Meeting
	mtg.getMeeting(name)
	os.Mkdir("out", 0777)
	makeImage(mtg)
	makePost(mtg)
	cmd := exec.Command("/bin/bash", "-c",
		fmt.Sprintf("rm -rf ../site/content/events/%[1]v && cp -pRP out ../site/content/events/%[1]v && rm -rf out", strings.Split(name, ".")[0]))
	cmd.Run()
}

func makeNotices() {
	yamlFile, err := ioutil.ReadFile("../site/data/events.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c Events
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	for _, v := range c.Meetings {
		if v.Date.After(time.Now()) {
			mtg := Meeting{
				Start:     v.Date.Add(time.Hour * 18),
				End:       v.Date.Add(time.Hour * 20),
				Location:  v.Location,
				Ed:        v.Education,
				Notes:     []string{"prospective members welcome"},
				Gauntlet:  "TBD",
				Title:     "General Meeting",
				PostTitle: v.Date.Format("January") + " General Meeting",
			}
			os.Mkdir("out", 0777)
			makeImage(mtg)
			makePost(mtg)
			cmd := exec.Command("/bin/bash", "-c",
				fmt.Sprintf("rm -rf ../site/content/events/%[1]v && cp -pRP out ../site/content/events/%[1]v && rm -rf out", v.Date.Format("2006-01")+"-gm.md"))
			cmd.Run()
		}
	}
}

func makePost(mtg Meeting) {
	mtg.Date = time.Now().Format("2006-01-02T15:04:05Z")
	mtg.Time = mtg.Start.Format("January 2 3pm")
	f, err := os.Create("out/index.md")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	defer writer.Flush()
	t := template.Must(template.New("event-template").ParseFiles("event-template.md"))
	err = t.ExecuteTemplate(writer, "event-template.md", mtg)
	if err != nil {
		panic(err)
	}
}

func makeImage(mtg Meeting) {
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	t := template.Must(template.New("magick-template").ParseFiles("magick-template.sh"))
	err := t.ExecuteTemplate(writer, "magick-template.sh", mtg)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	cmd := exec.Command("/bin/bash", "-c", b.String())
	cmd.Run()
}

func (c *Meeting) getMeeting(name string) *Meeting {

	yamlFile, err := ioutil.ReadFile(name)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

type Meeting struct {
	Start     time.Time `yaml:"start"`
	End       time.Time `yaml:"end"`
	Location  string    `yaml:"location"`
	Gauntlet  string    `yaml:"gauntlet"`
	Ed        string    `yaml:"ed"`
	Notes     []string  `yaml:"notes"`
	Title     string    `yaml:"title"`
	Date      string
	Time      string
	PostTitle string `yaml:"post-title"`
}

type Post struct {
	Title   string
	Content string
	Date    string
}

type Event struct {
	Location  string    `yaml:"location"`
	Info      string    `yaml:"info"`
	Date      time.Time `yaml:"date"`
	Url       string    `yaml:"url"`
	Education string    `yaml:"education"`
}
type Events struct {
	Meetings []Event `yaml:"meetings"`
	Club     []Event `yaml:"club"`
	Beer     []Event `yaml:"beer"`
}
