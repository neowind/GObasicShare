package slice

import (
	"fmt"
)

func Demo1() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s = %v\n", s)

	e := s
	fmt.Printf("e = %v | e := s\n", e)

	e = s[:]
	fmt.Printf("e = %v | e = s[:]\n", e)

	e = s[0:]
	fmt.Printf("e = %v | e = s[0:]\n", e)

	e = s[0:len(s)]
	fmt.Printf("e = %v | e = s[:len(s)]\n", e)

	e = s[1:9]
	fmt.Printf("e = %v | e = s[1:9]\n", e)
}

// func Demo2() {
// 	text := "0123456789ABCDEF"
// 	s := []byte(text)

// 	fmt.Println(s)
// }

func Demo3() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //Array
	fmt.Printf("s = %v\n", s)

	e := s
	fmt.Printf("Address[s] = %p\n", &s)
	fmt.Printf("Address[e] = %p\n", &e)

	s[9] = -1
	fmt.Printf("s = %v\n", s)
	fmt.Printf("e = %v\n", e)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SlictCap() {
	e := []int{}

	printSlice("e", e)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice("s", s)

	s2 := s[:0]
	printSlice("s2", s2)
	fmt.Println(s2)

	e = append(e, 0)
	printSlice("e", e)
}

func MyAppend() {
	e := []int{}
	s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	e = append(e, 0)
	printSlice("e", e)
	e = append(e, 1)
	printSlice("e", e)
	e = append(e, 2)
	printSlice("e", e)
	e = append(e, 3)
	printSlice("e", e)
	e = append(e, 4)
	printSlice("e", e)
	e = append(e, 5)
	printSlice("e", e)

	e = append(e, s...)
	printSlice("e + s", e)

	ss := [][]int{
		[]int{1, 2},
		[]int{11, 12},
		[]int{21, 22},
	}

	fmt.Printf("ss len=%d cap=%d %v\n", len(ss), cap(ss), ss)
}

func MyMake() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 3, 5)
	printSlice("b", b)

}
