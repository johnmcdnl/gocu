package gocu

import "github.com/cucumber/gherkin-go"

type Scenario struct {
	Pickle *gherkin.Pickle
	Timer *Timer
	Steps []*Step
}
