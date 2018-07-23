package main

import (
	"fmt"
	"sync"
	"time"

	"./KT/OOP"
	"./KT/routine"
)

func NewObjectEx() {
	//h := new(oop.Human)
	//h1 := oop.Human{}
	//h2 := &oop.Human{}

	// h1 := oop.Human{High: 180, Weight: 80}
	// h2 := &oop.Human{
	// 	High:   180,
	// 	Weight: 80,
	// }

	h3 := oop.NewHuman(180, 80)
	//fmt.Printf("h Type = %T \t Value = %v\n", h, h)
	// fmt.Printf("h1 Type = %T \t Value = %v\n", h1, h1)
	// fmt.Printf("h2 Type = %T \t Value = %v\n", h2, h2)
	//fmt.Printf("h3 Type = %T \t Value = %v\n", h3, h3)

	// High := h3.GetH()
	// fmt.Printf("creammy high = %f !!!!!!!", High)

	h3.SetH(200)
	High := h3.GetH()
	fmt.Printf("creammy high = %f !!!!!!!", High)

}

func PrintAvenger(a oop.Avenger) {
	fmt.Printf("Type = %T\n", a)
	//fmt.Println("Name = %s", a.Name)
	a.ShowSkill()
	a.ShowPower()
}

func RoutineDemo() {
	var wg sync.WaitGroup
	wg.Add(2)

	count := 5

	go func() {
		routine.Count("T1", 1, count)
		wg.Done()
	}()

	count = 4
	go func() {
		routine.Count("T2", 1, count)
		wg.Done()
	}()

	wg.Wait()
}

func x() {}
func ChannelDemo() {
	var wg sync.WaitGroup
	receiver := make(chan int, 10)
	endSign := make(chan struct{})
	count := 10

	wg.Add(2)

	//sender
	go func() {
		defer func() {
			fmt.Println("End")
			endSign <- struct{}{}
			wg.Done()
		}()

		for i := 0; i < count; i++ {
			receiver <- i
			fmt.Printf("%d ->\n", i)

		}
	}()

	//receiever
	go func() {
		defer func() {
			fmt.Println("Done")
			wg.Done()
		}()

		for {
			select {
			case i := <-receiver:
				fmt.Printf("<- %d\n", i)
				time.Sleep(time.Second)
			case <-endSign:
				close(receiver)
				close(endSign)
				return
			}
		}

	}()

	wg.Wait()
}

func main() {
	//pointer.Demo()
	//slice.Demo1()
	//slice.Demo2()
	//slice.Demo3()
	//slice.SlictCap()
	//slice.MyAppend()
	//slice.MyMake()
	//NewObjectEx()

	// ironman := oop.NewIronMan("Oni", "rish", 500, *oop.NewHuman(180, 80))
	// PrintAvenger(ironman)

	// w := oop.WarM{
	// 	Ironman:  ironman,
	// 	NickName: "Big O",
	// }

	// PrintAvenger(w)

	//RoutineDemo()
	ChannelDemo()

}
