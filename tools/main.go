package main

import (
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

var (
	outPath      = "../mainsrc/generated/out"
	paramPath    = "internal/test/params.yaml"
	templatePath = "internal/test/example_template"
)

func main() {
	err := writeTemplateWithParams()
	if err != nil {
		log.Fatal(err)
	}
}

func writeTemplateWithParams() error {
	outpath := "../mainsrc/generated/out"
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = fill(f)
	return err
}

type Params struct {
	Name     string `yaml:"name"`
	Gift     string `yaml:"gift"`
	Attended bool   `yaml:"attended"`
}

func fill(wr io.Writer) error {
	params := readParamFromYaml()
	err := fillTemplateWithParams(wr, params)
	if err != nil {
		return err
	}
	return nil
}

func readParamFromYaml() Params {
	data := mustReadFile(paramPath)
	var params Params
	err := yaml.Unmarshal([]byte(data), &params)
	if err != nil {
		log.Fatal("failed to parse yaml", err)
	}
	return params
}

func fillTemplateWithParams(wr io.Writer, params any) error {
	templateContents := mustReadFile(templatePath)
	err := template.Must(template.New("").Parse(templateContents)).Execute(wr, params)
	return err
}
