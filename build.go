package gocu

import (
	"github.com/cucumber/gherkin-go"
	"io/ioutil"
	"bytes"
)

func buildFeature(filePath string) (*Feature, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	gd, err := gherkin.ParseGherkinDocument(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	f, err := buildFeatureFromGherkinDocument(gd)
	f.Path = filePath
	if err != nil {
		return nil, err
	}
	return f, nil

}

func buildFeatureFromGherkinDocument(document *gherkin.GherkinDocument) (*Feature, error) {
	var f = &Feature{
		GherkinDocument: document,
		Timer:           &Timer{},
	}
	scenarios, err := buildScenarios(document.Pickles())
	if err != nil {
		return nil, err
	}
	f.Scenarios = scenarios
	return f, nil
}

func buildScenarios(pickles []*gherkin.Pickle) ([]*Scenario, error) {
	var scenarios []*Scenario
	for _, pickle := range pickles {
		scenario, err := buildScenario(pickle)
		if err != nil {
			return nil, err
		}
		scenarios = append(scenarios, scenario)
	}
	return scenarios, nil
}

func buildScenario(pickle *gherkin.Pickle) (*Scenario, error) {
	steps, err := buildSteps(pickle)
	if err != nil {
		return nil, err
	}

	var scenario = &Scenario{
		Timer:  &Timer{},
		Pickle: pickle,
		Steps:  steps,
	}

	return scenario, nil
}

func buildSteps(pickle *gherkin.Pickle) ([]*Step, error) {
	var steps []*Step
	for _, pickleStep := range pickle.Steps {
		step, err := buildStep(pickleStep)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}

func buildStep(pickleStep *gherkin.PickleStep) (*Step, error) {
	var step = &Step{
		Timer:      &Timer{},
		PickleStep: pickleStep,
	}

	return step, nil
}
