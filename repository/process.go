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

func (p *ListProcess) WorkingWithNode() (result float64, message string, err error) {
	count = 0

	var list []float64

	//Valida se a lista informada está vazia
	if len(p.List) == 0 {
		return 0, "A lista informada está vazia", errors.New("")
	}

	//Valida se os itens do array são numeros
	for _, valueList := range p.List {
		switch valueList := valueList.(type) {
		case float64:
			list = append(list, valueList)
		default:
			value, err := strconv.ParseFloat(valueList.(string), 64)
			if err != nil {
				return 0, fmt.Sprintf("Campo inválido: '%s'", valueList), err
			}

			list = append(list, value)
		}
	}

	//Busca nas variáveis do sistema a quantidade de 'nós'
	node, _ := strconv.Atoi(os.Getenv("NODE"))
	wg.Add(node)

	//Realiza a soma dos valores da lista, de acordo com a quantidade de nós
	for i := 1; i <= node; i++ {
		go Sum(list)
	}
	wg.Wait()

	return count, "Operação realizada com sucesso", nil
}

func Sum(list []float64) {
	defer wg.Done()

	for _, number := range list {
		count += number
	}
}
