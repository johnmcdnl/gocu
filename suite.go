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
	fn := func(path string, info os.FileInfo, err error) error{
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
		s.Features = append(s.Features, f)
		return nil
	}
	if err := filepath.Walk(s.Dir, fn); err!=nil{
		return err
	}
	return nil
}