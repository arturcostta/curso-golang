package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 1545.78,
		"Pedro junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 13250.00
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
