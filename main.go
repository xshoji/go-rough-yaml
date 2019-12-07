package main

import (
	"fmt"
	"github.com/xshoji/go-rough-yaml/goroughyaml"
)

func main() {
	roughYaml := goroughyaml.FromYaml(getSimpleYaml())

	roughYaml.Get("aaa").Get("aaa").Get("aaa").Get("aaa").Get("aaa").Get("aaa").SetForce("bbb", "ccc")

	fmt.Printf("print yaml :\n %v", roughYaml.ToYaml())
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
