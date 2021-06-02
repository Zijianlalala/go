package main
import (
	"fmt"
	"log"
	"github.com/headfirstgo/datafile"
)
func main()  {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, num := range numbers {
		sum += num
	}
	avg := sum / float64(len(numbers))
	fmt.Printf("average of numbers is %.2f\n", avg)
}