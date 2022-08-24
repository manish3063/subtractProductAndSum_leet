package main

import "fmt"

func main() {
	n := 234
	result := subtractProductAndSum(n)
	fmt.Println(result)
}
func subtractProductAndSum(n int) int {
	//fmt.Println(n)
	var res int = 0
	var sum int = 0
	var product int = 1

	for n > 0 {
		res = n % 10
		sum = sum + res
		product = product * res
		n = n / 10
	}
	final := product - sum
	//fmt.Println(final)
	return final

}
