package goroughyaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"testing"
)

func TestFromYaml(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
  ccc:
    - 1
`
	roughYaml := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (yaml.MapSlice)
	v1 := roughYaml.Get("aaa").Get("bbb").Get("bbb1").Value().(string)
	v2 := roughYaml.Get("aaa").Get("ccc").Get("0").Value().(int)
	if v1 != "bbb" || v2 != 1 {
		t.Errorf("<< FAILED >>> : orderedMapSlice.Get(\"aaa\").GetContents() is not yaml.MapSlice")
	}
	t.Logf("%v\n", roughYaml.GetContents())
}

func TestGetContents(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
    - 1
    - 2
    - c
`
	var actualValue interface{}
	var actualValuePtr *interface{}

	roughYamlObj := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (yaml.MapSlice)
	actualValue = roughYamlObj.Get("aaa").GetContents()
	_, ok := actualValue.(*yaml.MapSlice)
	if !ok {
		t.Errorf("<< FAILED >>> : roughYamlObj.Get(\"aaa\").GetContents() is not yaml.MapSlice")
	}
	t.Logf("%v\n", actualValue)

	//
	//
	//---------------------
	// success (slice)
	actualValuePtr, _ = roughYamlObj.Get("aaa").Get("ccc").GetContents().(*interface{})
	fmt.Printf("%v\n", reflect.TypeOf(*actualValuePtr).Kind())
	switch reflect.TypeOf(*actualValuePtr).Kind() {
	case reflect.Slice:
	default:
		t.Errorf("<< FAILED >>> : roughYamlObj.Get(\"aaa\").Get(\"ccc\").GetContents() is not []interface{}")
	}
	t.Logf("%v\n", roughYamlObj)
}

func TestGet(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
    - 1
    - 2
    - c
  ddd:
    - 3
    - 
      - 4
      - 5
`

	var expectedKey interface{}
	var expectedValue interface{}
	var expectedValueList []interface{}
	var actualKey interface{}
	var actualValue interface{}
	var actualValueList []interface{}

	orderedMapSlice := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (single value)
	expectedKey = "bbb1"
	expectedValue = "bbb"
	actualKey = orderedMapSlice.Get("aaa").Get("bbb").Get("bbb1").Key().(string)
	actualValue = orderedMapSlice.Get("aaa").Get("bbb").Get("bbb1").Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)

	//
	//
	//---------------------
	// success (single value)
	expectedKey = "bbb2"
	expectedValue = 111
	actualKey = orderedMapSlice.Get("aaa").Get("bbb").Get("bbb2").Key().(string)
	actualValue = orderedMapSlice.Get("aaa").Get("bbb").Get("bbb2").Value().(int)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)

	//
	//
	//---------------------
	// success (slice)
	expectedKey = "ccc"
	expectedValueList = []interface{}{1, 2, "c"}
	actualKey = orderedMapSlice.Get("aaa").Get("ccc").Key().(string)
	actualValueList = orderedMapSlice.Get("aaa").Get("ccc").Value().([]interface{})
	if actualKey != expectedKey || !compareSlice(actualValueList, expectedValueList) {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValueList:%v, expectedValueList:%v\n", actualKey, expectedKey, actualValueList, expectedValueList)

	//
	//
	//---------------------
	// success (value of slice)
	expectedKey = nil
	expectedValue = 2
	actualKey = orderedMapSlice.Get("aaa").Get("ccc").Get("1").Key()
	actualValue = orderedMapSlice.Get("aaa").Get("ccc").Get("1").Value().(int)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)

	//
	//
	//---------------------
	// success (nested slice)
	expectedKey = nil
	expectedValue = 5
	actualKey = orderedMapSlice.Get("aaa").Get("ddd").Get("1").Get("1").Key()
	actualValue = orderedMapSlice.Get("aaa").Get("ddd").Get("1").Get("1").Value().(int)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)

	//
	//
	//---------------------
	// success (nil)
	expectedKey = nil
	expectedValue = nil
	actualKey = orderedMapSlice.Get("xxx").Key()
	actualValue = orderedMapSlice.Get("xxx").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)

	//
	//
	//---------------------
	// success (nil)
	expectedKey = nil
	expectedValue = nil
	actualKey = orderedMapSlice.Get("xxx").Get("yyy").Get("zzz").Key()
	actualValue = orderedMapSlice.Get("xxx").Get("yyy").Get("zzz").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
}

func TestSet(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
  - 1
  - 2
`
	var expectedKey interface{}
	var expectedValue interface{}
	var actualKey interface{}
	var actualValue interface{}

	roughYamlObj := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (set value)
	expectedKey = "bbb1"
	expectedValue = "ccc"
	roughYamlObj.Get("aaa").Get("bbb").Set("bbb1", "ccc")
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: ccc
    bbb2: 111
  ccc:
  - 1
  - 2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set nil)
	expectedKey = "bbb1"
	expectedValue = nil
	roughYamlObj.Get("aaa").Get("bbb").Set("bbb1", nil)
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: null
    bbb2: 111
  ccc:
  - 1
  - 2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set nil 2)
	expectedKey = "bbb"
	expectedValue = nil
	roughYamlObj.Get("aaa").Set("bbb", nil)
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb: null
  ccc:
  - 1
  - 2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set MapSlice)
	expectedKey = "key-a"
	expectedValue = "value-a"
	value := yaml.MapSlice{
		yaml.MapItem{Key: "key-a", Value: "value-a"},
		yaml.MapItem{Key: "key-b", Value: "value-b"},
	}
	roughYamlObj.Get("aaa").Set("bbb", value)
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("key-a").Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Get("key-a").Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    key-a: value-a
    key-b: value-b
  ccc:
  - 1
  - 2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	////---------------------
	//// success (set slice value)
	//expectedKey = nil
	//expectedValue = 5
	//roughYamlObj.Get("aaa").Get("ccc").Set("1", 5)
	//actualKey = roughYamlObj.Get("aaa").Get("ccc").Get("1").Key()
	//actualValue = roughYamlObj.Get("aaa").Get("ccc").Get("1").Value().(int)
	//if actualKey != expectedKey || actualValue != expectedValue {
	//	t.Errorf("<< FAILED >>>")
	//}
	//t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	//bytes, _ = yaml.Marshal(roughYamlObj.GetContents())
	//fmt.Printf("---\n%v\n\n", string(bytes))
}

func TestSetForce(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
  - 1
  - 2
`
	var expectedKey interface{}
	var expectedValue interface{}
	var actualKey interface{}
	var actualValue interface{}

	roughYamlObj := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (set value)
	expectedKey = "ccc"
	expectedValue = "ccc-value1"
	roughYamlObj.Get("aaa").SetForce(expectedKey.(string), expectedValue.(string))
	actualKey = roughYamlObj.Get("aaa").Get(expectedKey.(string)).Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get(expectedKey.(string)).Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc: ccc-value1
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set new value)
	expectedKey = "ddd"
	expectedValue = "ddd-value1"
	roughYamlObj.Get("aaa").SetForce(expectedKey.(string), expectedValue.(string))
	actualKey = roughYamlObj.Get("aaa").Get(expectedKey.(string)).Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get(expectedKey.(string)).Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc: ccc-value1
  ddd: ddd-value1
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set new nested value)
	expectedKey = "ddd"
	expectedValue = "ddd-value1"
	roughYamlObj.Get("aaa").SetForce("111", nil)
	roughYamlObj.Get("aaa").Get("111").SetForce(expectedKey.(string), expectedValue.(string))
	actualKey = roughYamlObj.Get("aaa").Get("111").Get(expectedKey.(string)).Key().(string)
	actualValue = roughYamlObj.Get("aaa").Get("111").Get(expectedKey.(string)).Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc: ccc-value1
  ddd: ddd-value1
  "111":
    ddd: ddd-value1
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set to undefined key)
	expectedKey = nil
	expectedValue = nil
	roughYamlObj.Get("aaa").Get("aaa").Get("aaa").Get("aaa").SetForce("111", nil)
	roughYamlObj.Get("aaa").Get("aaa").Get("aaa").Get("aaa").Get("111").SetForce("ddd", "ddd-value")
	actualKey = roughYamlObj.Get("aaa").Get("aaa").Get("aaa").Get("aaa").Get("111").Get("ddd").Key()
	actualValue = roughYamlObj.Get("aaa").Get("aaa").Get("aaa").Get("aaa").Get("111").Get("ddd").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc: ccc-value1
  ddd: ddd-value1
  "111":
    ddd: ddd-value1
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}
}

func TestSetSlice(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
`
	var expectedKey interface{}
	var expectedValue interface{}
	var expectedValueList []interface{}
	var actualKey interface{}
	var actualValue interface{}
	var actualValueList []interface{}

	roughYamlObj := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (set slice)
	expectedKey = "bbb1"
	expectedValueList = []interface{}{"aaa", "bbb", "ccc"}
	roughYamlObj.Get("aaa").Get("bbb").Set("bbb1", []interface{}{"aaa", "bbb", "ccc"})
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Key().(string)
	actualValueList = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Value().([]interface{})
	if actualKey != expectedKey || !compareSlice(actualValueList, expectedValueList) {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValueList:%v, expectedValueList:%v\n", actualKey, expectedKey, actualValueList, expectedValueList)
	expectedValue = `aaa:
  bbb:
    bbb1:
    - aaa
    - bbb
    - ccc
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actual:%v, expected:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (set nested slice)
	expectedKey = "bbb1"
	expectedValueList = []interface{}{"aaa", "bbb", []interface{}{"ccc", "ddd", "eee"}}
	roughYamlObj.Get("aaa").Get("bbb").Set("bbb1", []interface{}{"aaa", "bbb", []interface{}{"ccc", "ddd", "eee"}})
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Key().(string)
	actualValueList = roughYamlObj.Get("aaa").Get("bbb").Get("bbb1").Value().([]interface{})
	if actualKey != expectedKey || !compareSlice(actualValueList, expectedValueList) {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValueList:%v, expectedValueList:%v\n", actualKey, expectedKey, actualValueList, expectedValueList)
	expectedValue = `aaa:
  bbb:
    bbb1:
    - aaa
    - bbb
    - - ccc
      - ddd
      - eee
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}
}

func TestDelete(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    ccc-key1: ccc-value1
    ccc-key2: ccc-value2
  ddd:
  - ddd1
  - ddd2
`

	var expectedKey interface{}
	var expectedValue interface{}
	var actualKey interface{}
	var actualValue interface{}

	roughYamlObj := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (delete key)
	expectedKey = nil
	expectedValue = nil
	roughYamlObj.Get("aaa").Get("bbb").Delete("ccc-key2")
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("ccc-key2").Key()
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Get("ccc-key2").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb:
    ccc-key1: ccc-value1
  ddd:
  - ddd1
  - ddd2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}

	//
	//
	//---------------------
	// success (delete key2)
	expectedKey = nil
	expectedValue = nil
	roughYamlObj.Get("aaa").Get("bbb").Delete("ccc-key1")
	actualKey = roughYamlObj.Get("aaa").Get("bbb").Get("ccc-key1").Key()
	actualValue = roughYamlObj.Get("aaa").Get("bbb").Get("ccc-key1").Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
	}
	t.Logf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	expectedValue = `aaa:
  bbb: null
  ddd:
  - ddd1
  - ddd2
`
	actualValue, _ = roughYamlObj.ToYaml()
	if actualValue != expectedValue {
		t.Errorf("<< FAILED >>>")
		t.Logf("actualValue:%v, expectedValue:%v\n", actualValue, expectedValue)
	}
}

func compareSlice(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		v := a[i]
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			compareSlice(toSlice(a[i]), toSlice(b[i]))
		default:
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}

func TestNext(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
    - '1'
    - '2'
    - c
  ddd:
    - 3
    - 
      - 4
      - 5
  eee:
    - 
      aaa: aaa1
      bbb: 
        - bbb1
        - bbb2
    - 
      aaa: aaa2
      bbb: 
        - bbb3
        - bbb4
`
	roughYaml := FromYaml(yamlString)

	//
	//
	//---------------------
	// success (yaml.MapSlice)
	v1 := roughYaml.Get("aaa").Get("ccc")
	for v1.HasNext() {
		n := v1.Next()
		v := n.Value().(string)
		if v != "1" && v != "2" && v != "c" {
			t.Errorf("<< FAILED >>> : roughYaml.Get(\"aaa\").Get(\"bbb\").Get(\"ccc\")")
		}
		t.Logf("%v\n", v)
	}

	v1 = roughYaml.Get("aaa").Get("ddd").Get("1")
	for v1.HasNext() {
		n := v1.Next()
		v := n.Value().(int)
		if v != 4 && v != 5 {
			t.Errorf("<< FAILED >>> : roughYaml.Get(\"aaa\").Get(\"bbb\").Get(\"ddd\").Get(\"1\")")
		}
		t.Logf("%v\n", v)
	}

	v1 = roughYaml.Get("aaa").Get("eee")
	for v1.HasNext() {
		v2 := v1.Next()
		v3 := v2.Get("bbb")
		for v3.HasNext() {
			n := v3.Next()
			v := n.Value().(string)
			if v != "bbb1" && v != "bbb2" && v != "bbb3" && v != "bbb4" {
				t.Errorf("<< FAILED >>> : roughYaml.Get(\"aaa\").Get(\"bbb\").Get(\"ddd\").Get(\"1\")")
			}
			t.Logf("%v\n", v)
		}
	}
}

func toSlice(value interface{}) []interface{} {
	slice := make([]interface{}, 0)
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		newValue := reflect.ValueOf(value)
		for i := 0; i < newValue.Len(); i++ {
			slice = append(slice, newValue.Index(i).Interface())
		}
	}
	return slice
}

func printYaml(roughYamlObj roughYaml) {
	yamlString, err := roughYamlObj.ToYaml()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("---\n%v\n\n", yamlString)
}
