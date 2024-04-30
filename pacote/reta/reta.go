package main

import "math"

// inicializando com letra maiuscula é Público(visível dentro e fora do pacote)
// inicializando com letra minúscula é PRIVADO(visível dentro do pacote)

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return //retorna o cx, cy por causa do retorno limpo
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
