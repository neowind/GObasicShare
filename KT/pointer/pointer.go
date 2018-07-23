package pointer

import "fmt"

func Demo() {
	i, j := 10, 20

	fmt.Printf("Address[i] = %p i = %d\n", &i, i)
	fmt.Printf("Address[j] = %p j = %d\n", &j, j)
	fmt.Println()

	fmt.Println("Point p to j")
	p := &i
	fmt.Printf("Type[p] %T\n", p)
	fmt.Printf("Address[p] = %p p = %d\n", p, *p)
	fmt.Println()

	*p = 11
	fmt.Println("Change p to 11")
	fmt.Printf("p = %d | i = %d\n", *p, i)
	fmt.Println()

	p = &j
	fmt.Println("Point p to j")
	fmt.Printf("Address[p] = %p p = %d\n", p, *p)
	fmt.Println()

	*p = 21
	fmt.Printf("p = %d | i = %d | j = %d\n", *p, i, j)
}
