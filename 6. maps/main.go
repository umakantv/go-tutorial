package main

import "fmt"

func main() {
	myMap := map[string]string{
		"umakant": "vashishtha",
		"kumar":   "sanu",
		"ajeet":   "kajala",
	}

	fmt.Println(myMap)

	// deleting entries from map
	delete(myMap, "ajeet")
	changeMap(myMap)

	// fmt.Println(myMap)

	printMap(myMap)

}

func printMap(myMap map[string]string) {
	// iterating over maps
	for name, surname := range myMap {
		fmt.Println(name, surname)
	}
}

// Since maps are Reference type data structures, it modifies the original passed value
func changeMap(myMap map[string]string) {
	myMap["dude"] = "chill"
}
