package main

import (
	"fmt"
	"runtime"
)

type Point struct {
	x int
	y int
}

func max(n1, n2 int) int {
	max := n1
	if n1 < n2 {
		max = n2
	}
	return max
}

func compute(num1, num2 int, add func(n, m int) int) int {
	return add(num1, num2)
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("%d: %d\n", i, sum)
	}

	sum = 0
	for sum < 1000 {
		sum += 1
	}
	fmt.Printf("sum: %d\n", sum)

	if sum == 1000 {
		fmt.Println("if statements don't need ()")
	}

	n1, n2 := 4, 6
	fmt.Printf("n1: %d, n2: %d, max: %d\n", n1, n2, max(n1, n2))

	// short if statements
	// variables declared inside short if statements will be
	// accessible inside any following else statements
	if num := 0; num > 0 {
		fmt.Println("num is greater than 0")
	} else if num < 0 {
		fmt.Println("num is greater than 0")
	} else {
		fmt.Println("num is equal to 0")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	defer fmt.Println("done")
	fmt.Printf("calculating max: %d\n", max(0, 1))

	num := 10
	p := &num
	fmt.Printf("p: %d, *p: %d\n", p, *p)

	point := Point{x: 1, y: 2}
	fmt.Println(point)
	point.x = 7
	point_ptr := &point
	(*point_ptr).y = 0 // same as point_ptr.y
	fmt.Println(*point_ptr)

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr[0:2])

	m := make(map[string]int)
	m["test1"] = 1
	m["test2"] = 5
	m["test3"] = 9
	m["test4"] = 420
	fmt.Println(m)
	delete(m, "test1")
	fmt.Println(m)
	val, present := m["test1"]
	fmt.Println("the value", val, "is present?", present)

	// FUNCTIONAL PROGRAMMING!
	fmt.Println(compute(5, 6, func(n, m int) int { return n + m }))
	fmt.Println(compute(5, 6, func(n, m int) int { return n * m }))
}
