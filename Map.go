package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//declare map
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("Map nil")
	} else {
		fmt.Println("Not nil")
	}

	fmt.Println()
	//add element to map
	myMap["Key1"] = 1
	myMap["Key2"] = 2
	fmt.Println(myMap)

	fmt.Println()
	//delete element on map
	delete(myMap, "key1")
	fmt.Println(myMap)

	//length map
	fmt.Println()
	fmt.Println(len(myMap))

	//map la reference type
	fmt.Println()
	myMap1 := myMap
	fmt.Println("new Map", myMap1)
	myMap1["Key2"] = 100

	fmt.Println("Map", myMap)
	fmt.Println("Map1", myMap1)

	//get element by key
	value, existed := myMap["Key1"] //no key existed
	if existed {
		fmt.Println("Co tim thay gia tri voi key 100")
	} else {
		fmt.Println("Khong tim thay gia tri voi key")
	}
	fmt.Println(value)

	//trong map khong co tuan tu == de so sanh map

	var rabbit = make(map[string]interface{})
	rabbit["id"] = 1
	rabbit["name"] = "Teo"
	rabbit["age"] = 2

	data, _ := json.Marshal(rabbit)
	var myJsonString = string(data)
	fmt.Println(myJsonString)
	fmt.Println(data)
}
