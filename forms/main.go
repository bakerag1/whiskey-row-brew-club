package main

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	paths := []string{
		"competition/festival-table-scores.tmpl",
		"competition/gauntlet-ballots.tmpl",
	}

	for _, v := range paths {
		writePdf(v)
	}
}

func writePdf(path string) {
	var png []byte
	url := "https://whiskeyrowbrewclub.com/forms"
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		QRCode string
		Url    string
	}{
		QRCode: base64.StdEncoding.EncodeToString(png),
		Url:    url,
	}

	basename := strings.Split(strings.Split(path, ".")[0], "/")[1]
	out := "output/" + basename + ".html"
	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	t := template.Must(template.New(basename).ParseFiles(path))
	err = t.ExecuteTemplate(writer, basename+".tmpl", data)
	if err != nil {
		panic(err)
	}
	writer.Flush()

	margin := "4"
	c := exec.Command("wkhtmltopdf", "-B", margin, "-L", margin, "-R", margin,
		out, "../site/content/forms/"+basename+".pdf")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
