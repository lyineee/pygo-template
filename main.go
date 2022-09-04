package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"
)

var tmplFile string
var outputFile string
var jsonDataFile string

func main() {
	flagInit()

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("opening template file '%s' fail, err: %s", tmplFile, err.Error())
	}

	// data
	var data interface{}
	if jsonDataFile != "" {
		f, err := os.Open(jsonDataFile)
		if err != nil {
			log.Printf("opening json data file '%s' fail, err: %s", jsonDataFile, err.Error())
		}
		defer f.Close()
		json.NewDecoder(f).Decode(&data)
		// log.Println(data)
	}

	// output
	if outputFile != "" {
		f, err := os.Create(outputFile)
		if err != nil {
			log.Printf("creating output file '%s' fail, err: %s", outputFile, err.Error())
		}
		tmpl.Execute(f, data)
	} else {
		tmpl.Execute(os.Stdout, data)
	}

}

func flagInit() {
	flag.StringVar(&tmplFile, "tmplFile", "template.tmpl", "template file path")
	flag.StringVar(&tmplFile, "f", "template.tmpl", "template file path")
	flag.StringVar(&outputFile, "output", "", "output file path")
	flag.StringVar(&outputFile, "o", "", "output file path")
	flag.StringVar(&jsonDataFile, "data", "", "data file path, json format")
	flag.StringVar(&jsonDataFile, "d", "", "data file path, json format")

	flag.Parse()
}
