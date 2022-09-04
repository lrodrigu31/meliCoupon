package services

import (
	"fmt"
	"math"
)

func CrearMatriz() {
	/*	peso := []int{
			0, 4, 3, 5, 2,
		}
		valor := []int{
			0, 10, 40, 30, 20,
		}
		pesoMaximo := 8
	*/

	peso := []int{
		0, 100, 210, 260, 80, 90,
	}
	valor := []int{
		0, 10, 10, 10, 10, 10,
	}
	pesoMaximo := 500

	filas := len(peso)
	columnas := pesoMaximo + 1

	var matrix [][]int

	for i := 0; i < filas; i++ {
		var vectorFila []int
		for j := 0; j < columnas; j++ {
			vectorFila = append(vectorFila, 0)
		}
		matrix = append(matrix, vectorFila)
	}

	//fmt.Println("matrix", matrix)

	for i := 1; i < len(peso); i++ {
		for j := 1; j <= pesoMaximo; j++ {
			if i == 1 {
				if j >= peso[i] {
					matrix[i][j] = valor[i]
				}
			} else if j < peso[i] {
				matrix[i][j] = matrix[i-1][j]
			} else {
				matrix[i][j] = int(math.Max(float64(matrix[i-1][j]), float64(valor[i]+matrix[i-1][j-peso[i]])))
			}
		}
	}

	//fmt.Println("matrix", matrix)

	var objeto []int

	j := pesoMaximo
	for i := len(peso) - 1; i > 0; i-- {
		if matrix[i][j] != matrix[i-1][j] && matrix[i][j] == (matrix[i-1][j-peso[i]]+valor[i]) {
			objeto = append(objeto, i)
			j -= peso[i]
		}
	}

	fmt.Println(objeto)
}
