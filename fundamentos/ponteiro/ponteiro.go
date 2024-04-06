package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereco da variável
	*p++   // desreferenciando (pegando valor)
	i++
	// Go nao tem aritmetica de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)

}
