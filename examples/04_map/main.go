package main

import "fmt"

func main() {
	// define map, key string and value string
	var strMap map[string]string
	// we can use make to initialize our map
	strMap = make(map[string]string)
	strMap["name"] = "Ady Rahmat Ma"
	strMap["job"] = "Software Engineer"

	fmt.Println(strMap)

	// or we could define it in a single command
	strMap = map[string]string{
		"name": "Ady Rahmat MA",
		"job":  "Software Engineer",
	}

	fmt.Println(strMap)

	// accessing map
	// play with strMap
	// because strMap is map[string]string, so name would be string
	name := strMap["name"]
	fmt.Println(name)

	// access undefined key, x would be a type of string but the value is ""
	x := strMap["x"]
	fmt.Println(x)

	// is there any better way? yeah, `exist` is boolean. returning true if exist
	y, exist := strMap["y"]
	if !exist {
		fmt.Printf("Pretty bad, y isn't exist. So y is '%s' or empty string\n", y)
	}

	// deleting key and value
	delete(strMap, "job")
	fmt.Println(strMap)

	// what if we need another type of value
	intMap := map[string]int{
		"age":    24,
		"weight": 59,
	}

	fmt.Println(intMap)

	// accessing map
	// play with intMap
	// because intMap is map[string]int, so age would be int
	age := intMap["age"]
	fmt.Println(age)

	// we could use a same method on checking some key is exist or not
	xy, exist := intMap["somekey"]
	if exist {
		// only print if it exist
		fmt.Println(xy)
	}

	// and what if we need a map with multiple type of value
	var inMap map[string]interface{}
	inMap = make(map[string]interface{})
	inMap["strMap"] = strMap
	inMap["intMap"] = intMap

	fmt.Println(inMap)

	// or we could use
	inMap = map[string]interface{}{
		"name": "Ady Rahmat MA",
		"age":  24,
	}

	fmt.Println(inMap)

	// accessing map
	// play with inMap
	// because inMap is map[string]interface{}, so we need type assertion to get the real value
	name, ok := inMap["name"].(string)
	if ok {
		// ok means, that the interface value of name has been transformed to string successfully
		fmt.Println(name)
	}

	age, ok = inMap["age"].(int)
	if ok {
		fmt.Println(age)
	}

	// but what if we access undefined key and doing type assertion?
	unknownKey, ok := inMap["somekey"].(float64)
	if !ok {
		fmt.Printf("somekey is not float64. unknownKey value is %v (default value of float64) \n", unknownKey)
	}

	// and then what if we access existing key, but doing the wrong type assertion?
	newName, ok := inMap["name"].(int)
	if !ok {
		fmt.Printf("name is not and int or didn't exist. newName value is empty string %s\n", newName)
	}

	// single command to define nested map
	nestedMap := map[string]interface{}{
		"personal": map[string]interface{}{
			"name": "Ady Rahmat MA",
			"age":  24,
			"job": map[string]string{
				"primary":   "Software Engineer",
				"secondary": "Woman's boyfriend",
			},
		},
	}
	fmt.Println(nestedMap)
}
