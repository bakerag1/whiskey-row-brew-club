package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"whiskeyrowbrewclub.com/styles"
)

func main() {
	for _, s := range styles.StylesBjcp2021 {
		writeStylePage(s)
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
