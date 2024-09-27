package main

// O range permite que você percorra um array. Sempre que é chamado, retorna dois valores: o índice e o valor.
func Soma(numeros []int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// SomaTudo calcula as respectivas somas de cada slice recebido
func SomaTodoOResto(numerosPararSomar ...[]int) []int {
	var somas []int

	for _, numeros := range numerosPararSomar {
		somas = append(somas, Soma(numeros))
	}

	return somas
}
