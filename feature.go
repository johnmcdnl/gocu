package gocu

import "github.com/cucumber/gherkin-go"

type Feature struct {
	Path string
	GherkinDocument *gherkin.GherkinDocument
	Timer *Timer
	Scenarios []*Scenario
}
