# go-rough-yaml

go-rough-yaml provides simple accessors to edit yaml.  
This means that you haven't to prepare a struct type.  
Additionally, go-rough-yaml preserves an order of map structure, so that when yaml is reverted to string, the keys of map are not sorted.

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
