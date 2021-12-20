package main

import (
	"fmt"
)

type Matrix struct {
	Valores      []float64
	Alto         int
	Ancho        int
	EsCuadratica bool
	Max          float64
}

func max(valores ...float64) float64 {
	max := valores[0]
	for _, val := range valores {
		if val > max {
			max = val
		}
	}
	return max
}

func (m *Matrix) Set(values ...float64) {
	m.Valores = append(m.Valores, values...)
	m.Alto = 3
	m.Ancho = 3
	m.EsCuadratica = true
	m.Max = 20
}

func (m Matrix) Print() {
	cont := 0
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf("%.2f ", m.Valores[cont])
			cont++
		}
		fmt.Println()
	}
}

func main() {
	var m Matrix
	m.Set(2.2, 6.7, 3.4, 22.6, 130.9, 6.7, 7.8, 9.0, 22.1)
	m.Print()
}
