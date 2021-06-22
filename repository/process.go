package repository

import (
	"errors"
	"fmt"
	"imports/models"
	"os"
	"strconv"
	"sync"
)

type ListProcess models.ListProcess

var count float64
var wg sync.WaitGroup

func (p *ListProcess) WorkingWithNode() (result float64, message string, nodeErr interface{}, err error) {
	count = 0

	//Busca nas variáveis do sistema a quantidade de 'nós'
	node, _ := strconv.Atoi(os.Getenv("NODE"))

	//Valida se a lista informada está vazia
	if len(p.List) == 0 {
		return 0, "A lista informada está vazia", nil, errors.New("")
	}

	listSplit := SplitList(p.List, node)

	listValid, message, nodeErr := IsValid(listSplit)
	if message != "" {
		return 0, message, nodeErr, errors.New("")
	}

	wg.Add(node)
	for _, list := range listValid {
		go Sum(list)
	}
	wg.Wait()

	return count, "Operação realizada com sucesso", nil, nil
}

//Sum realiza a soma dos valores da lista
func Sum(list []float64) {
	defer wg.Done()

	for _, number := range list {
		count += number
	}
}

//SplitList divide igualmente os valores da lista para cada 'nó'
func SplitList(list []interface{}, node int) [][]interface{} {

	listSplit := make([][]interface{}, node)

	iCont := 0

	for _, value := range list {
		listTemp := []interface{}{}

		if iCont >= node {
			iCont = 0
		}

		listTemp = append(listTemp, value)
		listSplit[iCont] = append(listSplit[iCont], listTemp...)

		iCont++
	}

	return listSplit
}

//IsValid valida se há algum valor informado na lista, que seja diferente de um numérico
func IsValid(listProcess [][]interface{}) (list [][]float64, message string, nodeErr interface{}) {
	listValid := make([][]float64, len(listProcess))

	for i, valueList := range listProcess {
		for _, valueSubList := range valueList {
			switch valueSubList := valueSubList.(type) {
			case float64:
				listValid[i] = append(listValid[i], valueSubList)
			default:
				value, err := strconv.ParseFloat(valueSubList.(string), 64)
				if err != nil {
					return nil, fmt.Sprintf("Campo inválido: '%s'", valueSubList), i + 1
				}

				listValid[i] = append(listValid[i], value)
			}
		}
	}

	return listValid, "", nil
}
