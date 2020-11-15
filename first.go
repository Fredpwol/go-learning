package main

import "fmt"

const (
	isAdmin = 1 << iota
	worksAtGoogle
	male
	female
	worksFromHome
	worksAtMicrosoft
	python
	golang
)

func main() {
	var chocies byte = isAdmin | male | python | worksFromHome
	// var age int = 20
	// var name string = "fred"
	// height := 180.2
	fmt.Printf("%b\n", chocies)
	fmt.Printf("isAdmin: %v\n", isAdmin&chocies == isAdmin)
	fmt.Printf("worksAtGoogle: %v\n", worksAtGoogle&chocies == worksAtGoogle)
	fmt.Printf("Male: %v\n", male&chocies == male)
	fmt.Printf("Female: %v\n", female&chocies == female)
	fmt.Printf("Works from Home: %v\n", worksFromHome&chocies == worksFromHome)
	fmt.Printf("Works At Microsoft %v\n", worksAtMicrosoft&chocies == worksAtMicrosoft)
	fmt.Printf("Codes in Python: %v\n", python&chocies == python)
	fmt.Printf("Codes in Golang: %v\n", golang&chocies == golang)

	names := [4]string{"fred", "sam", "jane", "foo"}
	ages := []int{12, 32, 14, 45}
	a := make([]int, 0, 100)
	a = append(a, ages...)
	fmt.Printf("Names %v\n", names)
	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("a values: %v\n", a)
}
