package main

import (
	"fmt"
	"github.com/xshoji/go-rough-yaml/goroughyaml"
	"gopkg.in/yaml.v2"
)

func main() {
	var yamlString = getSimpleYaml()
	mapSlice := &yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlString), mapSlice)
	roughYaml := goroughyaml.NewRoughYaml(mapSlice)

	bytes, _ := yaml.Marshal(roughYaml.GetContents())
	fmt.Printf("%v", string(bytes))
}

func getSimpleYaml() string {
	return `
# Development teams records
development-teams:
  team-a:
    pc-app-name1:
      id: 1001
    pc-app-name2:
      id: 1002
    ranks:
    - 100
    - 1000
`
}
