package main

import (
	"bootstrap-application/start/cryptografic"
	"bootstrap-application/start/decryption"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	looping := true

	fmt.Println("Olá, seja bem vindo ao meu algoritmo de criptografia!")

	for looping {
		fmt.Print("Digite a frase para ser criptografada:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		sentence := scanner.Text()

		fmt.Println()

		cryptografic.Criptografic(strings.ToUpper(sentence))

		fmt.Println("Frase Criptografada: ", cryptografic.Cryptograf)
		fmt.Println("Chave Secreta: ", cryptografic.Key)

		decryption.Decrypt(cryptografic.Cryptograf, cryptografic.Key)

		fmt.Println("Frase descriptografada:", strings.ToLower(decryption.Decryption))

		fmt.Println()

		var number int
		fmt.Println("Digite 1 para continuar ou 0 para sair")
		fmt.Scan(&number)

		if number == 0 {
			looping = false
		}
	}
}
