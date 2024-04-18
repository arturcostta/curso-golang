package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
func inc2(n *int) {
	// * é usado para desreferencciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {

	n := 1
	inc1(n)
	fmt.Println(n)

	// & é usado para obter o enderećo da variável
	inc2(&n)
	fmt.Println(n)
}
