package gocu

import "github.com/cucumber/gherkin-go"

type Feature struct {
	Path string
	GherkinDocument *gherkin.GherkinDocument
	Scenarios []*Scenario
}
