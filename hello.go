package main

import "fmt"
import "reflect"

func main() {

	var nome = "Paulo"
	var versao float32 = 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}
