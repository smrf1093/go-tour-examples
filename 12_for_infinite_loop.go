package main

import "fmt"

func main() {
	sum := 0

	for {
		sum += 1
		if sum > 10{
			break
		}
	}

	fmt.Println(sum)
}


