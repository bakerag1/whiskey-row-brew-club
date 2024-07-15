package main

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"os/exec"
	"text/template"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	paths := []string{
		"competition/festival-table-scores.tmpl",
	}

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
	f, err := os.Create("output/festival-table-scores.html")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	t := template.Must(template.New("festival-table-scores").ParseFiles(paths...))
	err = t.ExecuteTemplate(writer, "festival-table-scores.tmpl", data)
	if err != nil {
		panic(err)
	}
	writer.Flush()

	margin := "4"
	c := exec.Command("wkhtmltopdf", "-B", margin, "-L", margin, "-R", margin,
		"output/festival-table-scores.html", "../site/content/forms/festival-table-scores.pdf")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
