package main

import (
	"bufio"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
	"gopkg.in/yaml.v2"
)

func main() {
	var m members
	var data [][]string
	m.getMembers()
	for _, v := range m.Members {
		writeMember(v)
		url := fmt.Sprintf("https://whiskeyrowbrewclub.com/members/%v", v.Id)
		data = append(data, []string{v.Name, url})
	}
	file, err := os.Create("memberlist.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)
}

func writeMember(m Member) {
	paths := []string{"member.tmpl"}
	f, err := os.Create(fmt.Sprintf("../site/content/members/%v.md", m.Id))
	if err != nil {
		panic(err)
	}
	var png []byte
	url := fmt.Sprintf("https://whiskeyrowbrewclub.com/members/%v", m.Id)
	png, err = qrcode.Encode(url, qrcode.Medium, 256)
	m.QRCode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	t := template.Must(template.New("member").Funcs(nil).ParseFiles(paths...))
	err = t.ExecuteTemplate(writer, "member.tmpl", m)
	if err != nil {
		panic(err)
	}
}

type members struct {
	Members []Member `yaml:"members"`
}
type Member struct {
	Name   string `yaml:"name"`
	Id     string `yaml:"id"`
	Since  string `yaml:"since"`
	QRCode string `yaml:"qr"`
}

func (m *members) getMembers() *members {

	yamlFile, err := ioutil.ReadFile("../data/members.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return m
}
