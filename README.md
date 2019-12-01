# go-rough-yaml

go-rough-yaml provide accessors to edit yaml data.

```go
var yamlString = `
ddd:
  ccc:
    c: value-c
    a: value-a
  bbb:
  - 10
  - 5
aaa:
  zzz: value-zzz
  yyy: value-yyy
  xxx: value-xxx
`
mapSlice := &yaml.MapSlice{}
yaml.Unmarshal([]byte(yamlString), mapSlice)
roughYaml := goroughyaml.NewRoughYaml(mapSlice)

// get value
fmt.Println(roughYaml.Get("ddd").Get("ccc").Get("a")) // => value-a
fmt.Println(roughYaml.Get("ddd").Get("bbb").Get("1")) // => 5

// set value
roughYaml.Get("aaa").Set("yyy", nil)
fmt.Println(roughYaml.Get("aaa").Get("yyy")) // -> nil

// print as yaml
/**
ddd:
  ccc:
    c: value-c
    a: value-a
  bbb:
  - 10
  - 5
aaa:
  zzz: value-zzz
  yyy: null
  xxx: value-xxx
 */
bytes, _ := yaml.Marshal(roughYaml.GetContents())
fmt.Printf("%v", string(bytes))
```

### Features

- Simple interface
- Schema-less
- Preserving order of map structure
