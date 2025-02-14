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
	fmt.Println("Jogo da advinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar esse número. O número é um inteiro entre 0 e 100")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)

		case chuteInt == x:
			fmt.Printf(
				"Parabéns o número era: %d\n"+
					"Você acertou em %d tentativas\n"+
					"Essas foram as suas tentaivas: %v\n",
				x, i, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Infelizmente, você não acertou o número, que era %d\n"+
			"Você teve 10 tentativas. \n"+
			"Essas foram as suas tentaivas: %v\n",
		x, chutes,
	)
}
