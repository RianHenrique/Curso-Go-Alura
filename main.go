package main

import (
	"fmt"
	"projeto/contas"
)

func main() {

	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 123, NumeroConta: 12313231}
	contaDoRian := contas.ContaCorrente{Titular: "Rian", NumeroAgencia: 123, NumeroConta: 12313231}
	fmt.Println(contaDoRian.NumeroConta)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Jaelisson"
	fmt.Println(*contaDaCris)

	fmt.Println(contaDoRian.Sacar(10000))
	fmt.Println(contaDoRian)
	fmt.Println(contaDoRian.Depositar(120.543))
	fmt.Println(contaDoRian)

	fmt.Println(contaDoRian, "\n", contaDoGuilherme)
	fmt.Println(contaDoRian.Transferencia(121200, &contaDoGuilherme))
	fmt.Println(contaDoRian, "\n", contaDoGuilherme)

	contaDoRian.Depositar(-200)
	fmt.Println(contaDoRian)

}
