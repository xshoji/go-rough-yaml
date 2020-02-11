## go-rough-yaml

go-rough-yaml provides simple accessors to access and edit yaml.  
This means that you haven't to prepare a struct type.  
Additionally, go-rough-yaml preserves an order of map structure, so that when yaml is reverted to string, the keys of map are not sorted.

```go
// create RoughYaml
roughYaml := goroughyaml.FromYaml(`
ddd:
  ccc:
    c: value-c
  bbb:
  - 10
aaa:
  zzz: value-zzz
`)

// get value
roughYaml.
Get("ddd").
  Get("ccc").
    Get("c").Value() // => value-c

// set value
roughYaml.Get("aaa").Set("zzz", nil)
roughYaml.
  Get("aaa").
    Get("zzz").Value()) // -> nil

// delete key
roughYaml.Delete("aaa")
roughYaml.Get("aaa").Value()) // -> nil

// print as yaml
/**
ddd:
  ccc:
    c: value-c
  bbb:
  - 10
 */
roughYaml.ToYaml()
```

### Features

- Simple interface
- Schema-less
- Preserving order of map structure

## License

MIT