package gocu

import "github.com/cucumber/gherkin-go"

type Scenario struct {
	Pickle *gherkin.Pickle

	Name  string
	Steps []*Step
	Timer *Timer
}
