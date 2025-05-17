package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"text/template"

	"github.com/Code-Hex/go-wordwrap"
	"whiskeyrowbrewclub.com/styles"
)

func main() {
	if os.Args[1] == "posts" {
		writeStyleImages(
			[]string{"11A", "11B", "11C", "12A", "12C", "13A", "13B", "13C", "16A", "16B", "16C", "16D", "17A", "17B", "17D"},
		)
	} else {
		for _, s := range styles.StylesBjcp2021 {
			writeStylePage(s)
		}
	}
}

func writeStylePage(s styles.Style) {
	paths := []string{"stylepage.tmpl"}
	f, err := os.Create(fmt.Sprintf("../site/content/styles/bjcp/2021/%v-%v.md", s.Style_id, s.Name))
	if err != nil {
		panic(err)
	}

	styleWriter := bufio.NewWriter(f)
	defer styleWriter.Flush()

	t, err := template.New("stylepage.tmpl").Funcs(template.FuncMap{"hasPrefix": strings.HasPrefix}).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(styleWriter, s)
	if err != nil {
		panic(err)
	}
}

func writeStyleImages(ids []string) {
	for _, v := range styles.StylesBjcp2021 {
		if slices.Contains(ids, v.Style_id) {
			fmt.Printf("writing %s", v.Style_id)
			v.Overall_impression = wordwrap.WrapString(v.Overall_impression, 40)
			// fmt.Println(v.Overall_impression)
			makeImage(v)
		}
	}
}

func makeImage(s styles.Style) {
	b := new(bytes.Buffer)
	writer := bufio.NewWriter(b)
	t := template.Must(template.New("magick-template").ParseFiles("magick-template.sh"))
	err := t.ExecuteTemplate(writer, "magick-template.sh", s)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	cmd := exec.Command("/bin/bash", "-c", b.String())
	cmd.Run()
}
