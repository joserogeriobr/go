package main

import (
	. "fmt"
	"os"
)

var saldo float32
var deposito float32
var saque float32

func depositar(a float32) float32 {
	saldo = saldo + a
	return saldo
}

func sacar(a float32) float32 {
	saldo = saldo - a
	return saldo
}

func main() {

	Printf("\n\n\nSistema Banco Go\n\n")
	Printf("\n\n")
Loop:

	Printf("1 - Depositar\n")
	Printf("2 - Sacar\n")
	Printf("3 - Saldo\n")
	Printf("4 - Sair\n")
	var opcao int
	Printf("Escolha uma das opções\n")
	Scan(&opcao)
	switch opcao {
	case 1:
		Printf("Digite a quantia que irá depositar: \n\n")
		Scan(&deposito)
		saldo = depositar(deposito)
		Printf("Você depositou : %g e seu saldo é de: %g\n\n", deposito, saldo)
		goto Loop
	case 2:
		Printf("Digite a quantia que irá sacar: \n\n")
		Scan(&saque)
		if saldo >= saque {
			sacar(saque)
		} else {
			Printf("Saque maior do que a quantia disponível!\n\n")
			goto Loop
		}
		Printf("Você sacou : %g e seu saldo é de: %g\n\n", saque, saldo)
		goto Loop

	case 3:
		Printf("Seu saldo é de: %g\n\n", saldo)
		goto Loop
	case 4:
		os.Exit(3)
	default:
		Printf("Opção incorreta!\n\n")
		goto Loop
	}

}
