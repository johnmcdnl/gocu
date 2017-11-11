package gocu

import "github.com/cucumber/gherkin-go"

type Step struct {
	PickleStep *gherkin.PickleStep
	Timer      *Timer
}
