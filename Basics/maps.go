package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType

	// mapVariable = make(map[keyType]valueType)

	// using a Map literal
	// mapVariable := map[keyType]valueType{
	// 	"key1": value1,
	// 	"key2": value2,
	// }

	myMap := make(map[string]int)
	fmt.Println("Initial map:", myMap)

	myMap["key1"] = 10
	myMap["key2"] = 20
	myMap["code"] = 30
	fmt.Println(myMap)
	fmt.Println("key1:", myMap["key1"])
	fmt.Println("key2:", myMap["key2"])
	fmt.Println("code:", myMap["code"])
	fmt.Println("Updated map:", myMap)

	myMap["code"] = 18
	fmt.Println("Updated map:", myMap)

	delete(myMap, "key1")
	fmt.Println("Map after deletion:", myMap)

	myMap["key1"] = 17
	myMap["key2"] = 18
	fmt.Println("Map after adding new values:", myMap)

	// clear(myMap)
	fmt.Println("Map after clearing:", myMap)

	// value, unknownvalue := myMap["key1"]
	// fmt.Println("Value for 'key1':", value)
	// fmt.Println("Is 'key1' present?", unknownvalue)

	_, unknownvalue := myMap["key1"]
	// fmt.Println("Value for 'key1':", value)
	fmt.Println("Is 'key1' present?", unknownvalue)

	myMap2 := map[string]int{"a": 1, "b": 2}

	myMap3 := map[string]int{"a": 1, "b": 2}
	fmt.Println("myMap2:", myMap2)

	fmt.Println("myMap3:", myMap3)

	if maps.Equal(myMap2, myMap) {
		fmt.Println("myMap2 is equal to myMap")
	}

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 is equal to myMap3")
	}

	for k, v := range myMap3 {
		fmt.Println("Key:", k, "Value:", v)
	}

	for _, v := range myMap3 {
		fmt.Println("Value:", v)
	}

	_, ok := myMap3["a"]
	fmt.Println("Is 'a' present?", ok)

	if ok {
		fmt.Println("Value for 'a':", myMap3["a"])
	} else {
		fmt.Println("'a' is not present in the map")
	}

	// the zero value of a map is nil. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println("Value for 'key' in nil map:", val)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println("Value for 'key' after initialization:", myMap4["key"])

	fmt.Println("myMap length:", len(myMap))

	// multidimensional map

	myMap5 := make(map[string]map[string]string)

	myMap5["outerKey1"] = make(map[string]string)
	// myMap5["outerKey1"]["innerKey1"] = "value1"
	// myMap5["outerKey1"]["innerKey2"] = "value2"
	// fmt.Println("Multidimensional map:", myMap5)

	myMap5["map1"] = myMap4
	fmt.Println("Multidimensional map:", myMap5)
}
