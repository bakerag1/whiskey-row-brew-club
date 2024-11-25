package main

import (
	"log"
	"os"
	"sort"
	"time"

	"github.com/dustin/go-humanize"
	"gopkg.in/yaml.v2"
)

func main() {
	s := GetWinners("../site/data/competitions.yaml")
	data, err := yaml.Marshal(&s)
	if err != nil {
		log.Fatal(err)
	}
	file := "../site/data/standings.yaml"
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
}

type Standing struct {
	Place  string
	Name   string
	Points int
}

type Comp struct {
	Winners []Winner    `yaml:"winners"`
	Type    string      `yaml:"compType"`
	Info    string      `yaml:"info"`
	Link    string      `yaml:"link"`
	Date    string      `yaml:"date"`
	Style   interface{} `yaml:"style"`
	Entries int         `yaml:"entries"`
}
type Schedules struct {
	Schedules []Schedule `yaml:"schedules"`
	Winners   []Winner   `yaml:"winners"`
	CompTypes []CompType `yaml:"compTypes"`
}
type Schedule struct {
	Year  string `yaml:"year"`
	Comps []Comp `yaml:"comps"`
}

type Winner struct {
	Name     string `yaml:"name"`
	Points   int    `yaml:"points"`
	Category string `yaml:"category"`
	Position string `yaml:"position"`
	Year     string `yaml:"year"`
	Post     string `yaml:"post"`
}

type CompType struct {
	Id          string         `yaml:"id"`
	DisplayName string         `yaml:"displayName"`
	Description string         `yaml:"description"`
	Points      map[string]int `yaml:"points"`
}

func GetWinners(file string) []Standing {

	var scheds Schedules
	yamlFile, err := os.Open(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	d := yaml.NewDecoder(yamlFile)
	d.SetStrict(true)
	err = d.Decode(&scheds)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	var thisYearsSched Schedule
	for _, v := range scheds.Schedules {
		if v.Year == time.Now().Format("2006") {
			thisYearsSched = v
			break
		}
	}
	var winners []Winner
	for _, v := range thisYearsSched.Comps {
		v = addPoints(v, scheds.CompTypes)
		if v.Winners != nil {
			winners = append(winners, v.Winners...)
		}
	}
	sTemp := CalcStandings(winners)
	var s []Standing
	for key, v := range sTemp {
		s = append(s, Standing{
			Place:  "",
			Name:   key,
			Points: v,
		})
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].Points == s[j].Points {
			return s[i].Name > s[j].Name
		}
		return s[i].Points > s[j].Points
	})
	place := 1
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i].Points != s[i-1].Points {
			place++
		}
		s[i].Place = humanize.Ordinal(place)
	}

	return s
}

func addPoints(v Comp, comps []CompType) Comp {
	for _, c := range comps {
		if v.Type == c.Id {
			for i, w := range v.Winners {
				v.Winners[i].Points = c.Points[w.Position]
			}
			break
		}
	}
	return v
}

func CalcStandings(w []Winner) map[string]int {
	standings := make(map[string]int)
	for _, v := range w {
		standings[v.Name] += v.Points
	}
	log.Printf("%v", standings)
	return standings
}
