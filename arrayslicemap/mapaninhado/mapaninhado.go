package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8456.78,
		},
		"J": {
			"Jose Joao": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)

		/*
			for funcionario, salario := range funcs {
				fmt.Println(funcionario, salario)
			}*/
	}

}
