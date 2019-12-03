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
	return NewRoughYaml(mapSlice)
}

func NewRoughYaml(yamlData interface{}) roughYaml {
	rootMapItem := yaml.MapItem{Key: "root", Value: yamlData}
	orderedMapSlice := roughYaml{
		contents:    yamlData,
		currentItem: &rootMapItem,
	}
	return orderedMapSlice
}

func createRoughYaml(item *yaml.MapItem, yamlData interface{}) *roughYaml {
	return &roughYaml{
		contents:    yamlData,
		currentItem: item,
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
	if o.contents == nil {
		return createRoughYaml(nil, nil)
	}
	mapSlice, ok := o.GetContents().(*yaml.MapSlice)
	//fmt.Printf("-- o.contents.(yaml.MapSlice)\n")
	//fmt.Printf("---- mapSlice	: %T, %p, %v\n", mapSlice, mapSlice, mapSlice)
	//fmt.Printf("---- ok : %v\n", ok)
	if ok {
		for index, item := range *mapSlice {
			referencedItem := &(*mapSlice)[index]
			//fmt.Printf("---- item.Value: %T, item: %p, key: %v, value: %v, value-pointer: %p, value-pointers pointer: %v\n", referencedItem.Value, referencedItem, referencedItem.Key, referencedItem.Value, referencedItem.Value, &referencedItem.Value)
			if referencedItem.Key == key {
				v, ok := referencedItem.Value.(yaml.MapSlice)
				if ok {
					return createRoughYaml(referencedItem, &v)
				}
				return createRoughYaml(referencedItem, &item.Value)
			}
		}
	}
	slice, ok := o.GetContents().(*interface{})
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
					return createRoughYaml(&v, &v.Value)
				}
			}
		}

	}
	return createRoughYaml(nil, nil)
}

func (o *roughYaml) Set(key string, value interface{}) {
	orderedMapSlice := o.Get(key)
	if orderedMapSlice == nil {
		return
	}
	orderedMapSlice.currentItem.Value = value
}

func (o *roughYaml) SetSlice(key string, slice []interface{}) {
	orderedMapSlice := o.Get(key)
	if orderedMapSlice == nil {
		return
	}
	orderedMapSlice.currentItem.Value = slice
}

func dumpNode(name string, node interface{}) {
	bytes, _ := yaml.Marshal(node)
	fmt.Printf("=======\n%v:\n%v=======\n", name, string(bytes))
}

func printPointer(name string, v interface{}, p interface{}) {
	fmt.Printf("-- %v => v %T: &v=%p v=&i=%p p=%p\n", name, v, &v, v, p)
}
