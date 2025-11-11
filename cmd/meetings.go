/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// meetingsCmd represents the meetings command
var meetingsCmd = &cobra.Command{
	Use:   "meetings",
	Short: "Add meeting occurrences",
	Long:  `given the year, create the club meetings`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := strconv.Atoi(args[0])
		fmt.Printf("wrbc meetings called for: %v\n", year)
		os.Mkdir("site/content/events/"+strconv.Itoa(year), os.ModePerm)
		for i := 1; i < 12; i++ {
			// skip july and december
			if i == 7 {
				continue
			}
			createMeeting(i, year)
		}
	},
}

func createMeeting(month int, year int) {
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	daysUntilThursday := (int(time.Thursday) - int(firstOfMonth.Weekday()) + 7) % 7
	thirdThursdayDay := 1 + daysUntilThursday + 14
	dt := time.Date(year, time.Month(month), thirdThursdayDay, 18, 0, 0, 0, time.Local)

	fmt.Printf("creating for month %v -> %s\n", month, dt.Format("2006-01-02 (Mon)"))
	dur, _ := time.ParseDuration("2h")
	data := struct {
		Month string
		Now   string
		Start string
		End   string
	}{
		Month: dt.Format("January"),
		Start: dt.Format(time.RFC3339),
		End:   dt.Add(dur).Format(time.RFC3339),
		Now:   time.Now().Format(time.RFC3339),
	}
	path := fmt.Sprintf("site/content/events/%v/%v-gm.md", strconv.Itoa(year), dt.Format("2006-01"))
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	template.Must(template.New("mtg").Parse(mtgTemplate)).Execute(writer, data)
}

func init() {
	rootCmd.AddCommand(meetingsCmd)
}

const mtgTemplate = `---
title: {{ .Month }} General Meeting
date: {{ .Now }}
startDate: {{ .Start }}
endDate:  {{ .End }}
location: TBD
event_type: meeting
meeting:
  ed: TBD
event: true
---
  
Notes:  
  
  * prospective members welcome  `
