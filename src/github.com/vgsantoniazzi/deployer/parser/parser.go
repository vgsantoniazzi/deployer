package parser

import (
  "io/ioutil"
  "path/filepath"
  "gopkg.in/yaml.v2"
)

type Auth struct {
  Username string
  Password string
  Host string
  Port string
}

type Run struct {
  Commands []string
  Auth Auth
}

func Access(filename string)(*Auth){
  run := Parse(filename)
  return &run.Auth
}

func Commands(filename string) ([]string){
  run := Parse(filename)
  return run.Commands;
}

func Parse(filename string) (Run) {
  file, _ := filepath.Abs(filename)
  yamlFile, err := ioutil.ReadFile(file)

  if err != nil {
    panic(err)
  }

  var run Run
  err = yaml.Unmarshal(yamlFile, &run)
  if err != nil {
    panic(err)
  }
  return run;
}
