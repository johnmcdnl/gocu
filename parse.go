package gocu

import (
	"github.com/cucumber/gherkin-go"
	"io/ioutil"
	"bytes"
)

func ParseDir(dir string) {

}

func ParseFile(filePath string) (*gherkin.GherkinDocument, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return gherkin.ParseGherkinDocument(bytes.NewReader(b))
}
