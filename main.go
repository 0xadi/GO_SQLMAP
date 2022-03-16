// package main

// func main() {
// 	a := get_body("http://localhost:8000/index.php?id=1")
// 	println(a)
// }
package main

import "fmt"

func main() {
	// a := get_body("http://localhost:8000/index.php?id=1")
	// fmt.Println(a)
	a := count_database("http://localhost:8000/index.php?id=1")
	b := database_payload("http://localhost:8000/index.php?id=1", a)
	fmt.Println(b)
}
