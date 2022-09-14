package main

import (
	"fmt"
	"time"
)

func main() {
	//These code is for calculate the valence electron of the an element
	var valence_electron_num, atomic_num int
	fmt.Printf("Input the Atomic number of the element: ")
	fmt.Scanln(&atomic_num)
	if atomic_num == 1 || atomic_num == 2 {
		fmt.Printf("The valence electron of the element is %d", atomic_num)
	} else {
		valence_electron_num = (atomic_num - 2) % 8
		if valence_electron_num == 0 {
			valence_electron_num = 8
		}
		fmt.Printf("The valence electron of the element is %d", valence_electron_num)
	}
	time.Sleep(5 * time.Second)
}
