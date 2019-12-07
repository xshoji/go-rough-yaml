// Package goroughyaml provides simple accessors to access and edit yaml.
// This means that you haven't to prepare a struct type.
// Additionally, go-rough-yaml preserves an order of map structure, so that when yaml is reverted to string, the keys of map are not sorted.
//
// Features
//
// - Simple interface
// - Schema-less
// - Preserving an order of map structure
//
// How to use
//
// Create object
//
//    roughYaml := goroughyaml.FromYaml(`
//    ddd:
//      ccc:
//        c: value-c
//        a: value-a
//      bbb:
//      - 10
//      - 5
//    aaa:
//      zzz: value-zzz
//      yyy: value-yyy
//      xxx: value-xxx
//    `)
//
//
// Get value
//
//    roughYaml.
//    Get("ddd").
//      Get("ccc").
//        Get("a").Value()) // => value-a
//
//    roughYaml.
//    Get("ddd").
//      Get("bbb").
//        Get("1").Value()) // => 5
//
//    roughYaml.Get("xxx").Value()) // => nil
//
//
// Set value
//
//    roughYaml.Get("aaa").Set("yyy", nil)
//    roughYaml.Get("aaa").Get("yyy").Value()) // -> nil
//
//
// Add value
//
//    roughYaml.Get("aaa").SetForce("bbb", "value-bbb")
//    roughYaml.Get("aaa").Get("bbb").Value()) // -> "value-bbb"
//
//
// Delete key
//
//    roughYaml.Delete("ddd")
//    roughYaml.Get("ddd").Value()) // -> nil
//    /**
//      aaa:
//        zzz: value-zzz
//        yyy: null
//        xxx: value-xxx
//    */
//
// Print as yaml
//
//    roughYaml.ToYaml()
//    /**
//      ddd:
//        ccc:
//          c: value-c
//          a: value-a
//        bbb:
//        - 10
//        - 5
//      aaa:
//        zzz: value-zzz
//        yyy: null
//        xxx: value-xxx
//    */
//
// Source code and other details for the project are available at GitHub:
//
//   https://github.com/xshoji/go-rough-yaml
//
package goroughyaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"strconv"
)

type roughYaml struct {
	contents    interface{}
	currentItem *yaml.MapItem
}

func FromYaml(yamlContent string) roughYaml {
	mapSlice := &yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlContent), mapSlice)
	return newRoughYaml(mapSlice)
}

func newRoughYaml(yamlData interface{}) roughYaml {
	rootMapItem := yaml.MapItem{Key: "root", Value: yamlData}
	orderedMapSlice := roughYaml{
		contents:    yamlData,
		currentItem: &rootMapItem,
	}
	return orderedMapSlice
}

func createRoughYaml(yamlContents interface{}, item *yaml.MapItem) *roughYaml {
	return &roughYaml{
		contents:    yamlContents,
		currentItem: item,
	}
}

func createRoughYamlNil() *roughYaml {
	return &roughYaml{
		contents:    nil,
		currentItem: nil,
	}
}

func (o *roughYaml) ToYaml() string {
	bytes, _ := yaml.Marshal(o.GetContents())
	return string(bytes)
}

func (o *roughYaml) GetContents() interface{} {
	if o.contents == nil {
		return nil
	}
	v, ok := o.contents.(*yaml.MapSlice)
	if ok {
		return v
	}
	return o.contents
}

func (o *roughYaml) Key() interface{} {
	if o.currentItem != nil {
		return o.currentItem.Key
	}
	return nil
}

func (o *roughYaml) Value() interface{} {
	if o.currentItem != nil {
		return o.currentItem.Value
	}
	if o.contents != nil {
		return o.contents
	}
	return nil
}

func (o *roughYaml) Get(key string) *roughYaml {
	//dumpNode("o.GetContents", o.GetContents())
	//fmt.Printf(">> o.contents : %T, %v\n", o.contents, o.contents)
	contents := o.GetContents()
	if contents == nil {
		return createRoughYamlNil()
	}
	mapSlice, ok := contents.(*yaml.MapSlice)
	//fmt.Printf("-- o.contents.(yaml.MapSlice)\n")
	//fmt.Printf("---- mapSlice	: %T, %p, %v\n", mapSlice, mapSlice, mapSlice)
	//fmt.Printf("---- ok : %v\n", ok)
	if ok {
		for index := range *mapSlice {
			referencedItem := &(*mapSlice)[index]
			//fmt.Printf("---- item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", referencedItem.Value, referencedItem, referencedItem.Key, referencedItem.Value, referencedItem.Value, &referencedItem.Value)
			if referencedItem.Key == key {
				if referencedItem.Value == nil {
					return createRoughYaml(nil, referencedItem)
				}
				mapSlicePointer, ok := referencedItem.Value.(*yaml.MapSlice)
				if ok {
					return createRoughYaml(mapSlicePointer, referencedItem)
				}
				mapSliceValue, ok := referencedItem.Value.(yaml.MapSlice)
				if ok {
					return createRoughYaml(&mapSliceValue, referencedItem)
				}
				return createRoughYaml(&referencedItem.Value, referencedItem)
			}
		}
	}
	slice, ok := contents.(*interface{})
	//fmt.Printf("--o.contents.(*interface{})\n")
	//fmt.Printf("---- slice : %T, %v\n", slice, slice)
	//fmt.Printf("---- ok : %v\n", ok)
	if ok {
		// > go - range over interface{} which stores a slice - Stack Overflow
		// > https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice?answertab=active#tab-top
		switch reflect.TypeOf(*slice).Kind() {
		case reflect.Slice:
			//fmt.Printf(">> slice is slice!\n")
			s := reflect.ValueOf(*slice)
			for i := 0; i < s.Len(); i++ {
				index := strconv.FormatInt(int64(i), 10)
				//fmt.Printf("---- index : %v, key : %v\n", index, key)
				//fmt.Printf("---- index == key : %v\n", index == key)
				if index == key {
					//fmt.Printf("---- item: %T\n", s.Index(i).Interface())
					v := yaml.MapItem{Key: nil, Value: s.Index(i).Interface()}
					return createRoughYaml(&v.Value, &v)
				}
			}
		}
	}
	return createRoughYamlNil()
}

func (o *roughYaml) Set(key string, value interface{}) {
	o.setValue("value", false, key, value, nil)
}

func (o *roughYaml) SetSlice(key string, slice []interface{}) {
	o.setValue("slice", false, key, nil, slice)
}

func (o *roughYaml) SetForce(key string, value interface{}) {
	o.setValue("value", true, key, value, nil)
}

func (o *roughYaml) SetSliceForce(key string, slice []interface{}) {
	o.setValue("slice", true, key, nil, slice)
}

func (o *roughYaml) setValue(valueType string, isForce bool, key string, value interface{}, slice []interface{}) {
	orderedMapSlice := o.Get(key)
	if orderedMapSlice.currentItem == nil {
		if isForce == false {
			return
		}
		if o.GetContents() == nil {
			return
		}
		content := o.GetContents()
		mapSlice, ok := content.(*yaml.MapSlice)
		newMapSlice := yaml.MapSlice{}
		newMapItem := yaml.MapItem{
			Key:   key,
			Value: nil,
		}
		if ok {
			for index := range *mapSlice {
				referencedItem := &(*mapSlice)[index]
				newMapSlice = append(newMapSlice, *referencedItem)
			}
		}
		newMapSlice = append(newMapSlice, newMapItem)
		setContentsValue(o, &newMapSlice)
		orderedMapSlice = o.Get(key)
	}

	if valueType == "value" {
		setContentsValue(orderedMapSlice, value)
	} else if valueType == "slice" {
		setContentsSlice(orderedMapSlice, slice)
	}
}

func setContentsValue(o *roughYaml, value interface{}) {
	o.contents = value
	o.currentItem.Value = value
}

func setContentsSlice(o *roughYaml, slice []interface{}) {
	o.contents = slice
	o.currentItem.Value = slice
}

func (o *roughYaml) Delete(key string) {
	if o.contents == nil {
		return
	}
	mapSlice, ok := o.GetContents().(*yaml.MapSlice)
	newMapSlice := yaml.MapSlice{}
	if ok {
		for index, _ := range *mapSlice {
			referencedItem := &(*mapSlice)[index]
			//fmt.Printf("---- item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", referencedItem.Value, referencedItem, referencedItem.Key, referencedItem.Value, referencedItem.Value, &referencedItem.Value)
			if referencedItem.Key != key {
				newMapSlice = append(newMapSlice, *referencedItem)
			}
		}
	}
	//fmt.Printf("---- newMapSlice: %T, newMapSlice-p: %p, newMapSlice-v: %v\n", newMapSlice, newMapSlice, newMapSlice)
	if len(newMapSlice) == 0 {
		setContentsValue(o, nil)
		return
	}

	setContentsValue(o, newMapSlice)
}

func dumpNode(name string, node interface{}) {
	bytes, _ := yaml.Marshal(node)
	fmt.Printf("=======\n%v:\n%v=======\n", name, string(bytes))
}

func printPointer(name string, v interface{}, p interface{}) {
	fmt.Printf("-- %v => v %T: &v=%p v=&i=%p p=%p\n", name, v, &v, v, p)
}
