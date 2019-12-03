package main

import (
	"fmt"
	"github.com/xshoji/go-rough-yaml/goroughyaml"
)

func main() {
	roughYaml := goroughyaml.FromYaml(getSimpleYaml())

	fmt.Printf("development-teams.team-a.pc-app-name1.id : %v\n",
		roughYaml.Get("development-teams").
			/*   */ Get("team-a").
			/*     */ Get("pc-app-name1").
			/*       */ Get("id").Value())

	fmt.Printf("development-teams.team-a.ranks[0] : %v\n",
		roughYaml.Get("development-teams").
			/*   */ Get("team-a").
			/*     */ Get("ranks").
			/*       */ Get("0").Value())

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
