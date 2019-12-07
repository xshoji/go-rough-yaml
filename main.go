package main

import (
	"fmt"
	"github.com/xshoji/go-rough-yaml/goroughyaml"
)

func main() {

	roughYaml := goroughyaml.FromYaml(``)
	roughYaml.SetForce("zzz", nil)
	roughYaml.Get("zzz").SetForce("ccc", "ccc-value1")
	roughYaml.Get("zzz").SetForce("bbb", "bbb-value2")
	roughYaml.Get("zzz").SetForce("aaa", "aaa-value3")
	roughYaml.Get("zzz").SetForce("test", []interface{}{"a", "a", "a"})
	roughYaml.SetForce("yyy", nil)
	roughYaml.Get("yyy").SetForce("ccc", "ccc-value1")
	roughYaml.Get("yyy").SetForce("bbb", "bbb-value2")
	roughYaml.Get("yyy").SetForce("aaa", "aaa-value3")
	roughYaml.Get("yyy").SetForce("test2", []interface{}{"a", "a", "a"})

	fmt.Printf("print yaml :\n %v", roughYaml.ToYaml())

	roughYaml.Delete("zzz")

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
