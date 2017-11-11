package gocu

import (
	"github.com/cucumber/gherkin-go"
	"io/ioutil"
	"bytes"
)

func ParseFile(filePath string) (*Feature, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	gd, err := gherkin.ParseGherkinDocument(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	f, err := featureFromGherkinDoc(gd)
	f.Path = filePath
	if err != nil {
		return nil, err
	}
	return f, nil

}

func featureFromGherkinDoc(document *gherkin.GherkinDocument) (*Feature, error) {
	var f = &Feature{
		GherkinDocument: document,
	}
	return f, nil
}
