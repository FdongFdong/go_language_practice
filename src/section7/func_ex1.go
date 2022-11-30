package main
import "fmt"

func multiply(n ...int)int {
	tot := 1
	for _, value := range n{
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
		tot := 0
	for _, value := range n{
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 : ",value)
	}
}

func main() {
	x := multiply(1,2,3)
  y := sum(1,2,3,4,5,6,7,8,9,10)

	fmt.Println("ex1 : ",x)
	fmt.Println("ex1 : ",y)
	fmt.Println()

	prtWord("a","apple", "test","golang","seoul")

	// 예제 3
	a := []int{1,2,3,4,5}
	m:= multiply(a...)
	n:= sum(a...)

	fmt.Println("ex3: ",m)
	fmt.Println("ex3: ",n)

}