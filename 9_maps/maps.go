package main

import "fmt"

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

	// p1 := map[string]int{"price": 50, "phone": 3}
	// p2 := map[string]int{"price": 50, "phone": 8}
	// fmt.Println(maps.Equal(p1, p2))

	// m := map[string]string{"fname": "Govind", "lname": "Tiwari"}

	// // for k, val := range m {
	// // 	fmt.Println(k,val)
	// // }

	// // if u didn't give second return then it retutn key not value
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// in strings
	for i, c := range "Govind" {

		// unicode code point rune
		// i mean starting byte of rune
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}
