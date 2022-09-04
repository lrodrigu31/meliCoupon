package services

import (
	"coupon/models"
	"coupon/repository"
	"fmt"
)

func Calculate(items map[string]float64, amount float64) []string {

	for i, item := range items {
		fmt.Println("i:", i)
		fmt.Println("items:", item)
	}

	list := []string{"MLA1", "MLA2", "MLA4", "MLA5"}
	return list
}

func GetPriceList(itemIds []string) (models.PriceList, float64) {
	priceList := make(map[string]float64)
	curPrice := 0.0
	for _, itemId := range itemIds {
		item := repository.GetItem(itemId)
		if _, present := priceList[item.Id]; !present {
			if item.ValidateStructure() {
				minPrice(&curPrice, &item.Price)
				priceList[item.Id] = item.Price
			}
		}
	}
	listMapped := models.PriceList{}
	listMapped.Items = priceList
	return listMapped, curPrice
}

func minPrice(price *float64, priceToCompare *float64) {
	if *price == 0.0 {
		*price = *priceToCompare
	} else if *price > *priceToCompare {
		*price = *priceToCompare
	}
}

type Objetos struct {
	Items []Objeto
}

type Objeto struct {
	Peso      float64
	Beneficio int
}

type Solution struct {
	Solutions []int
}

func Algoritmo() {
	capacidad := 500.00         // pesoFinal de la mochila
	objetos := []models.Objeto{ //objetos que van o no en la mochila
		{260},
		{210},
		{100},
		{90},
		{80},
	}

	var solution []int

	for i := 0; i < len(objetos); i++ {
		solution = append(solution, -1)
	}
	//fmt.Println(solutionActual)
	pesoActual := 0.0

	_, _, solucionFinal, pesoFinal := models.MochilaRec(solution, 0, objetos, solution, pesoActual, capacidad, 5)
	fmt.Println(solucionFinal, pesoFinal)

}

func MochilaRec(solution []int, etapa int, objetos []Objeto, mochilaFinal []int, pesoFinal float64, beneficioFinal *int, capacidad float64) {
	i := 0

	fmt.Println(":: solution[etapa] :: ", solution[etapa])

	if etapa > (len(objetos) - 1) {
		return
	}

	for {
		//do {
		fmt.Println(":: i :: ", i)
		solution[etapa] = i
		if validarPeso(solution, etapa, objetos, capacidad) {
			if etapa == (len(objetos) - 1) {
				actualizarSolution(solution, objetos, mochilaFinal, pesoFinal, beneficioFinal)
			} else {
				MochilaRec(solution, etapa+1, objetos, mochilaFinal, pesoFinal, beneficioFinal, capacidad)
			}
		}
		i++
		//} while
		if solution[etapa] != -1 {
			break
		}
	}
	solution[etapa] = -1

}

func validarPeso(solution []int, etapa int, objeto []Objeto, capacidad float64) bool {

	sum := 0.0
	for i := 0; i < etapa; i++ {
		if solution[i] == 1 {
			sum += objeto[i].Peso
		}
	}
	return sum < capacidad
}

func actualizarSolution(solution []int, objetos []Objeto, mochilaFinal []int, pesoFinal float64, beneficioFinal *int) {
	var beneficioTotal int
	pesoTotal := 0.0

	for i := 0; i < len(objetos); i++ {
		if solution[i] == 1 {
		}
		beneficioTotal += objetos[i].Beneficio
		pesoTotal += objetos[i].Peso

	}
	if pesoTotal > pesoFinal {
		for i := 0; i < len(objetos); i++ {
			mochilaFinal[i] = solution[i]
		}
		*beneficioFinal = beneficioTotal
		pesoFinal = pesoTotal
	}

}

func validarPeso2(solution *Solution, etapa int, objeto []Objeto, capacidad float64) bool {

	sum := 0.0
	for i := 0; i < etapa; i++ {
		if solution.Solutions[i] == 1 {
			sum += objeto[i].Peso
		}
	}
	return sum < capacidad
}

func MochilaRec2(solution *Solution, etapa *int, objetos []Objeto, mochilaFinal *Solution, pesoFinal *float64, beneficioFinal *int, capacidad float64) (Solution, float64, int, Solution) {
	i := 0
	if *etapa > (len(objetos) - 1) {
		return *mochilaFinal, *pesoFinal, *beneficioFinal, *solution
	}

	for {
		//do {
		fmt.Println("etapa", *etapa)
		solution.Solutions[*etapa] = i
		fmt.Println("solution[etapa]", solution.Solutions[*etapa])
		if validarPeso2(solution, *etapa, objetos, capacidad) {

			if *etapa == (len(objetos) - 1) {
				*mochilaFinal, *pesoFinal, *beneficioFinal = actualizarSolution2(solution, objetos, mochilaFinal, pesoFinal, beneficioFinal)
			} else {
				nuevaEtapa := *etapa + 1
				*mochilaFinal, *pesoFinal, *beneficioFinal, *solution = MochilaRec2(solution, &nuevaEtapa, objetos, mochilaFinal, pesoFinal, beneficioFinal, capacidad)
			} /**/
		}
		i++
		fmt.Println("solution[etapa]", solution.Solutions[*etapa])
		//} while
		if solution.Solutions[*etapa] != 1 {
			break
		}
	}
	solution.Solutions[*etapa] = -1
	return *mochilaFinal, *pesoFinal, *beneficioFinal, *solution
}

func actualizarSolution2(solution *Solution, objetos []Objeto, mochilaFinal *Solution, pesoFinal *float64, beneficioFinal *int) (Solution, float64, int) {
	var beneficioTotal int
	pesoTotal := 0.0

	for i := 0; i < len(objetos); i++ {
		if solution.Solutions[i] == 1 {
			beneficioTotal += objetos[i].Beneficio
			pesoTotal += objetos[i].Peso
		}
	}
	if pesoTotal > *pesoFinal {
		for i := 0; i < len(objetos); i++ {
			mochilaFinal.Solutions[i] = solution.Solutions[i]
		}
		*beneficioFinal = beneficioTotal
		*pesoFinal = pesoTotal
	}
	return *mochilaFinal, *pesoFinal, *beneficioFinal

}
