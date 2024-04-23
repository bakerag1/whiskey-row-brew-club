package styles

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed bjcp_styleguide-2021.json
var styleguide2021 []byte

var StylesBjcp2021 = parseStyles()

func parseStyles() []Style {
	var styles StyleHolder
	err := json.Unmarshal(styleguide2021, &styles)
	if err != nil {
		log.Fatal(err)
	}
	return styles.Beerjson.Styles
}

type StyleHolder struct {
	Beerjson Beerjson `json:"beerjson"`
}

type Beerjson struct {
	Styles []Style `json:"styles"`
}

type Style struct {
	Name                           string `json:"name"`
	Category                       string `json:"category"`
	Category_id                    string `json:"category_id"`
	Style_id                       string `json:"style_id"`
	Category_description           string `json:"category_description"`
	Overall_impression             string `json:"overall_impression"`
	Aroma                          string `json:"aroma"`
	Appearance                     string `json:"appearance"`
	Flavor                         string `json:"flavor"`
	Mouthfeel                      string `json:"mouthfeel"`
	Comments                       string `json:"comments"`
	History                        string `json:"history"`
	Style_comparison               string `json:"style_comparison"`
	Tags                           string `json:"tags"`
	Original_gravity               MeasurementRange
	International_bitterness_units MeasurementRange
	Final_gravity                  MeasurementRange
	Alcohol_by_volume              MeasurementRange
	Color                          MeasurementRange
	Ingredients                    string `json:"ingredients"`
	Examples                       string `json:"examples"`
	Style_guide                    string `json:"style_guide"`
}

type Measurement struct {
	Unit  string
	Value float64
}

type MeasurementRange struct {
	Minimum Measurement
	Maximum Measurement
}
