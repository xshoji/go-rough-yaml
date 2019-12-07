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
//    roughYaml.Get("aaa").SetForce("ggg", "value-bbb")
//    roughYaml.Get("aaa").Get("ggg").Value()) // -> "value-ggg"
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
// Copyright
//
//   Copyright (c) 2019 xshoji
//   This software is released under the MIT License.
//   http://opensource.org/licenses/mit-license.php
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
	contents := o.GetContents()
	if contents == nil {
		return createRoughYamlNil()
	}
	mapSlice, ok := contents.(*yaml.MapSlice)
	if ok {
		for index := range *mapSlice {
			referencedItem := &(*mapSlice)[index]
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
	if ok {
		// > go - range over interface{} which stores a slice - Stack Overflow
		// > https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice?answertab=active#tab-top
		switch reflect.TypeOf(*slice).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(*slice)
			for i := 0; i < s.Len(); i++ {
				index := strconv.FormatInt(int64(i), 10)
				if index == key {
					v := yaml.MapItem{Key: nil, Value: s.Index(i).Interface()}
					return createRoughYaml(&v.Value, &v)
				}
			}
		}
	}
	return createRoughYamlNil()
}

func (o *roughYaml) Set(key string, value interface{}) {
	o.setValue(key, value, false)
}

func (o *roughYaml) SetForce(key string, value interface{}) {
	o.setValue(key, value, true)
}

func (o *roughYaml) setValue(key string, value interface{}, isForce bool) {
	childMapSlice := o.Get(key)
	if childMapSlice.currentItem == nil {
		if isForce == false {
			return
		}
		newMapSlice := yaml.MapSlice{}
		newMapItem := yaml.MapItem{
			Key:   key,
			Value: nil,
		}
		content := o.GetContents()
		mapSlice := &yaml.MapSlice{}
		ok := true
		if content != nil {
			mapSlice, ok = content.(*yaml.MapSlice)
		}
		if ok {
			for index := range *mapSlice {
				referencedItem := &(*mapSlice)[index]
				newMapSlice = append(newMapSlice, *referencedItem)
			}
		}
		newMapSlice = append(newMapSlice, newMapItem)
		setContentsValue(o, &newMapSlice)
		childMapSlice = o.Get(key)
	}

	setContentsValue(childMapSlice, value)
}

func setContentsValue(o *roughYaml, value interface{}) {
	if o.currentItem == nil {
		return
	}
	o.contents = value
	o.currentItem.Value = value
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
			if referencedItem.Key != key {
				newMapSlice = append(newMapSlice, *referencedItem)
			}
		}
	}

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
