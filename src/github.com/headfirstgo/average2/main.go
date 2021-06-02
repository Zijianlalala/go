package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
 )
func main()  {
	
	numbers := os.Args[1:]
	var sum float64 = 0
	for _, num := range numbers {
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	}
	avg := sum / float64(len(numbers))
	fmt.Printf("average of numbers is %.2f\n", avg)
}