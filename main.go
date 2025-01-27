package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Jogo da adivinhação\n")
	fmt.Printf("Escolha um número entre 1 e 100\n" +
		"Você tem 10 tentativas\n")

	scanner := bufio.NewScanner(os.Stdin)

	attemps := [10]int64{}

	x := rand.Int64N(100) + 1

	for i := range attemps {
		fmt.Println("Digite um número:")
		scanner.Scan()
		attemp := scanner.Text()
		attemp = strings.TrimSpace(attemp)

		attempInteger, err := strconv.ParseInt(attemp, 10, 64)

		if err != nil {
			fmt.Println("Número inválido")
			return
		}

		switch {
		case attempInteger < x:
			fmt.Println("O número é maior")
		case attempInteger > x:
			fmt.Println("O número é menor")
		case attempInteger == x:
			fmt.Printf("Parabéns, você acertou! O número é %d \n"+
				"Você acertou em %d tentativas \n"+
				"Aqui está suas tentativas: %v \n", x, i+1, attemps[:i])
			return
		}

		attemps[i] = attempInteger
	}

	fmt.Printf("Você perdeu! O número era %d\n", x)
	fmt.Println("Tentativas: ", attemps)

}
