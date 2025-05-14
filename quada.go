package main

import "fmt"

func QuadA(x,y int) {
	if x <= 0 || y <= 0{
		return
	} else {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 || i == y-1{
					if j == 0 || j == x-1 {
						fmt.Print("o")
					} else {
						fmt.Print("-")
					}
				} else {
					if j == 0 || j == x-1 {
						fmt.Print("|")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	QuadA(5,3)
}
