package goarea

import "math"

// PI proposção numérica
const PI = 3.1416

// Circ é responsavel por caluclar a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
