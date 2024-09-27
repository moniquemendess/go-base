package main

// O range permite que você percorra um array. Sempre que é chamado, retorna dois valores: o índice e o valor.
func Soma(numeros []int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
