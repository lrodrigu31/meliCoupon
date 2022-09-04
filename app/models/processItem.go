package models

type Objeto struct {
	Peso float64
}

func MochilaRec(solution []int, level int, objetos []Objeto, finalSolution []int, totalAmount float64, AMOUNT float64, N int) ([]int, int, []int, float64) {

	if level > N-1 {
		return solution, level, finalSolution, totalAmount
	}
	i := 0
	for {
		//do {
		solution[level] = i
		isValid := validarPeso(solution, level, objetos, AMOUNT)
		if isValid {
			if level == (N - 1) {
				finalSolution, totalAmount = actualizarSolution(solution, objetos, finalSolution, totalAmount, N)
			} else {
				solution, level, finalSolution, totalAmount = MochilaRec(solution, level+1, objetos, finalSolution, totalAmount, AMOUNT, N)
			}
		}
		i++
		//} while
		if solution[level] != -1 {
			break
		}
	}
	solution[level] = -1
	return solution, level, finalSolution, totalAmount
}

func actualizarSolution(solution []int, objetos []Objeto, mochilaFinal []int, pesoFinal float64, N int) ([]int, float64) {
	pesoTotal := 0.0
	for i, objeto := range objetos {
		if i < N {
			if solution[i] == 1 {
				pesoTotal += objeto.Peso
			}
		}
	}
	if pesoTotal >= pesoFinal {
		for i, itemSolution := range solution {
			if itemSolution == 1 {
				mochilaFinal[i] = itemSolution
			}

		}
		pesoFinal = pesoTotal
	}

	return mochilaFinal, pesoFinal
}

func validarPeso(solution []int, level int, objeto []Objeto, capacidad float64) bool {

	sum := 0.0
	for i, obj := range objeto {
		if i < level {
			if solution[i] == 1 {
				sum += obj.Peso
			}
		}
	}
	if sum < capacidad {
		return true
	}
	return false
}
