package gocu

import (
	"path/filepath"
	"os"
	"strings"
)

type Suite struct {
	Dir      string
	Features []*Feature
}

func NewSuite(dir string) (*Suite, error){
	var s = &Suite{
		Dir: dir,
	}
	if err:= s.Build(); err !=nil{
		return nil, err
	}
	return s ,nil
}

func (s *Suite) Build() error {
	parseGherkinFunc := func(path string, info os.FileInfo, err error) error{
		if info.IsDir(){
			return nil
		}
		if !strings.HasSuffix(path, ".feature"){
			return nil
		}
		f, err := ParseFile(path)
		if err!=nil{
			return err
		}
		if f.GherkinDocument.Feature ==nil{
			return nil
		}
		s.Features = append(s.Features, f)
		return nil
	}
	return filepath.Walk(s.Dir, parseGherkinFunc)
}