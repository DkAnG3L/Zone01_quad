package main

import "fmt"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				if j == 0 {
					fmt.Print("/")
				} else if j == x-1 {
					fmt.Print("\\")
				} else {
					fmt.Print("*")
				}
			} else if i == y-1 {
				if j == 0 {
					fmt.Print("\\")
				} else if j == x-1 {
					fmt.Print("/")
				} else {
					fmt.Print("*")
				}
			} else {
				if j == 0 || j == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
func main() {
	QuadB(5,3)
}