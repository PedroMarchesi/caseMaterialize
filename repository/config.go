package repository

import (
	"os"
	"strconv"
)

func CreatKnot(value string) (message string, err error) {

	valueNode, err := strconv.Atoi(value)

	//Caso err != nil, a variável contem um valor diferente de um número, impossibilitando a conclusão do processo.
	//Caso valueKnot seja menor que 0, a variável contém um valor negativo,  impossibilitando a conclusão do processo.
	if err != nil || valueNode < 0 {
		return "O valor informado para configurar a quantidade de nós está incorreto", err
	}

	//É usado o SetEnv para salvar a quantidade de "nós" em uma variavém de ambiente
	os.Setenv("KNOT", strconv.Itoa(valueNode))

	return "", nil
}
