# go-rough-yaml

go-rough-yaml provide accessors to edit yaml data.

```go
// create RoughYaml
roughYaml := goroughyaml.FromYaml(`
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
`)

// get value
fmt.Printf("%v\n", roughYaml.
Get("ddd").
  Get("ccc").
    Get("a").Value()) // => value-a
fmt.Printf("%v\n", roughYaml.
Get("ddd").
  Get("bbb").
    Get("1").Value()) // => 5

// set value
roughYaml.Get("aaa").Set("yyy", nil)
fmt.Printf("%v\n", roughYaml.
  Get("aaa").
    Get("yyy").Value()) // -> nil

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
fmt.Printf("%v", roughYaml.ToYaml())
```

### Features

- Simple interface
- Schema-less
- Preserving order of map structure
