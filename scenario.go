package gocu

import "github.com/cucumber/gherkin-go"

type Scenario struct {
	pickle *gherkin.Pickle
	Timer *Timer
	Steps []*Step
}
