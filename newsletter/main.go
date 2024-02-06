package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	yyyyMM := os.Args[1]
	paths := []string{
		"newsletter.tmpl",
	}

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
		"modIsZero": func(i int, m int) bool {
			return i%m == 0 && i != 0
		},
	}

	var cfg config
	cfg.getConf()

	f, err := os.Create(fmt.Sprintf("%v-newsletter.html", yyyyMM))
	if err != nil {
		panic(err)
	}
	var monthInfo monthData
	monthInfo.getMonthInfo(yyyyMM)
	tData := struct {
		Config    config
		MonthInfo monthData
	}{
		Config:    cfg,
		MonthInfo: monthInfo,
	}
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	t := template.Must(template.New("newsletter").Funcs(funcMap).ParseFiles(paths...))
	err = t.ExecuteTemplate(writer, "newsletter.tmpl", tData)
	if err != nil {
		panic(err)
	}
}

func (c *config) getConf() *config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *monthData) getMonthInfo(yyyyMM string) *monthData {

	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%v.yaml", yyyyMM))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	d, err := time.Parse("2006-January", yyyyMM)
	if err != nil {
		panic(err)
	}
	c.Month = d.Format("January")
	c.Year = d.Format("2006")
	return c
}

type link struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}
type config struct {
	Toplinks []link `yaml:"top-links"`
}
type monthData struct {
	ClubEvents   []event `yaml:"club-events"`
	BeerEvents   []event `yaml:"beer-events"`
	Competitions []event `yaml:"competitions"`
	Meetings     []event `yaml:"meetings"`
	News         string  `yaml:"news"`
	Board        []board `yaml:"board"`
	Month        string
	Year         string
}
type event struct {
	Name      string `yaml:"name"`
	Url       string `yaml:"url"`
	Info      string `yaml:"info"`
	Date      string `yaml:"date"`
	Gauntlet  string `yaml:"gauntlet"`
	Education string `yaml:"education"`
	Location  string `yaml:"location"`
}
type board struct {
	Position string `yaml:"position"`
	Name     string `yaml:"name"`
}
