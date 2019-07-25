package binary

import "fmt"

func main() {
	fmt.Println("Decimal - Binary")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d - %b\n", i, i)
	}
}
