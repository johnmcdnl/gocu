package main

import (
	"fmt"

	"github.com/johnmcdnl/gocu"
)

const testDir = `../../vendor/github.com/cucumber/gherkin-go/testdata/good`

func main() {
	fmt.Println("gocu")
	s, err := gocu.NewSuite(testDir)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	for _, f := range s.Features {
		fmt.Println(f.Path)
		fmt.Println(f.GherkinDocument.Feature.Name)
		for _, scenario := range f.Scenarios {
			for _, step := range scenario.Steps {
				fmt.Println(step.PickleStep.Text)
			}
		}
	}
}
