package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13)) // => 55
}

// if err := http.ListenAndServe(":8080", nil); err != nil {
// log.Fatal("ListenAndServe:", err)
// errはhttp.Li~~の返値でしかない
// }
