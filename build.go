package gocu

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cucumber/gherkin-go"
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
	if err != nil {
		return nil, fmt.Errorf("%s has error: %s", filePath, err.Error())
	}
	f.Path = filePath

	return f, nil

}

func buildFeatureFromGherkinDocument(document *gherkin.GherkinDocument) (*Feature, error) {
	scenarios, err := buildScenarios(document.Pickles())
	if err != nil {
		return nil, err
	}

	if document.Feature == nil {
		return nil, errors.New("feature tag not found")
	}

	if document.Feature.Name == "" {
		return nil, errors.New("feature does not have a name")
	}

	var f = &Feature{
		GherkinDocument: document,
		Name:            document.Feature.Name,
		Scenarios:       scenarios,
		Timer:           &Timer{},
	}
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
		Pickle: pickle,

		Name:  pickle.Name,
		Steps: steps,
		Timer: &Timer{},
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
		PickleStep: pickleStep,

		Name:  pickleStep.Text,
		Timer: &Timer{},
	}

	return step, nil
}
