package main

import (
	"fmt"
	"maps"
)

// maps => js obj, python dic., has maps
func main() {
	// m := make(map[string]string)
	// fmt.Println(m)

	// m["name"] = "GoLang"
	// fmt.Println(m)

	// fmt.Println(m["name"])

	// if key dosn't exits in the map it return zerod value according to val type
	// fmt.Println(m["phone"])

	// len for length

	// delete(m, "name")
	// clear(m)
	// fmt.Println(m)

	//without make func
	// p := map[string]int{"price": 50, "phone": 3}
	// fmt.Println(p)

	// // how element is in the map or not
	// v, ok := p["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("All OK")
	// } else {
	// 	fmt.Println("Not OK")
	// }

	p1 := map[string]int{"price": 50, "phone": 3}
	p2 := map[string]int{"price": 50, "phone": 8}
	fmt.Println(maps.Equal(p1, p2))

}
