package main

import "fmt"

func main() {
	//var aprovados map[int]string maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678913] = "Pedro"
	aprovados[22345678978] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])

}
