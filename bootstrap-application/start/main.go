package main

import (
	"bootstrap-application/start/cryptografic"
	"bootstrap-application/start/decryption"
	"fmt"
	"strings"
)

func main() {
	var sentence string
	looping := true

	fmt.Println("Olá, seja bem vindo ao meu algoritmo de criptografia!")
	for looping {
		fmt.Println("Digite a frase para ser criptografada:")
		fmt.Scan(&sentence)

		fmt.Println()
		cryptografic.Criptografic(strings.ToUpper(sentence))

		fmt.Println("Frase Criptografada: ")
		fmt.Println(cryptografic.Cryptograf)

		fmt.Println("Chave Secreta:")
		fmt.Println(cryptografic.Key)

		decryption.Decrypt(cryptografic.Cryptograf, cryptografic.Key)

		fmt.Println("Frase descriptografada:")
		fmt.Println(strings.ToLower(decryption.Decryption))

		fmt.Println()

		var number int
		fmt.Println("Digite 1 para continuar ou 0 para sair")
		fmt.Scan(&number)

		if number == 0 {
			looping = false
		}
	}
}
