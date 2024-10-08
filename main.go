package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	jogador    = "X"
	computador = "O"
	vazio      = "\u2000"
	tamanho    = 3
)

var tabuleiro [tamanho][tamanho]string
var gabarito [tamanho][tamanho]int

func main() {
	inicializaGabarito()
	inicializaTabuleiro()
	for {
		movimentoHumano()
		if verificaFimDoJogo(tabuleiro, jogador) {
			imprimeTabuleiroEGabarito()
			fmt.Println("Você ganhou!")
			break
		}
		if empate() {
			imprimeTabuleiroEGabarito()
			fmt.Println("Empate!")
			break
		}
		movimentoComputador()
		if verificaFimDoJogo(tabuleiro, computador) {
			imprimeTabuleiroEGabarito()
			fmt.Println("O computador ganhou!")
			break
		}
		if empate() {
			imprimeTabuleiroEGabarito()
			fmt.Println("Empate!")
			break
		}
	}
}

func inicializaTabuleiro() {
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			tabuleiro[i][j] = vazio
		}
	}
}

func inicializaGabarito() {
	contador := 1
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			gabarito[i][j] = contador
			contador++
		}
	}
}

func imprimeTabuleiroEGabarito() {
	fmt.Println("Tabuleiro        Gabarito")
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			if j < tamanho-1 {
				fmt.Printf("%s | ", tabuleiro[i][j])
			} else {
				fmt.Printf("%s", tabuleiro[i][j])
			}
		}

		fmt.Printf("        ")

		for j := 0; j < tamanho; j++ {
			if j < tamanho-1 {
				fmt.Printf("%d | ", gabarito[i][j])
			} else {
				fmt.Printf("%d", gabarito[i][j])
			}
		}
		fmt.Println()
	}
}

// Verifica se o valor inserido é um int ou um numero valido
func movimentoHumano() {
	leitor := bufio.NewReader(os.Stdin)
	var movimento int
	for {
		imprimeTabuleiroEGabarito()
		fmt.Println("Digite o número da casa que quer jogar: ")
		input, err := leitor.ReadString('\n')
		if err != nil {
			fmt.Println("Erro na leitura", err)
		}
		input = strings.TrimSpace(input)
		movimento, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Digite um numero valido.")
			continue
		}
		if movimento >= 1 && movimento <= 9 {
			linha := (movimento - 1) / 3
			coluna := (movimento - 1) % 3
			if tabuleiro[linha][coluna] == vazio {
				tabuleiro[linha][coluna] = jogador
				break
			} else {
				fmt.Println("Movimento inválido!")
			}
		}

	}
}

func movimentoComputador() {
	_, movimento := minimax(tabuleiro, 0, true)
	if movimento != [2]int{-1, -1} {
		tabuleiro[movimento[0]][movimento[1]] = computador
	}
}

// Algoritmo minimax
func minimax(tabuleiro [tamanho][tamanho]string, profundidade int, maximizar bool) (int, [2]int) {
	// Verifica se o computador venceu e retorna um valor positivo, quanto menor a profundidade, maior o valor.
	if verificaFimDoJogo(tabuleiro, computador) {
		return 10 - profundidade, [2]int{-1, -1}
	}

	// Verifica se o jogador venceu e retorna um valor negativo, quanto menor a profundidade, menor o valor.
	if verificaFimDoJogo(tabuleiro, jogador) {
		return profundidade - 10, [2]int{-1, -1}
	}

	// Verifica se houve empate, retornando 0.
	if empate() {
		return 0, [2]int{-1, -1}
	}

	var melhorValor int
	var melhorMovimento [2]int = [2]int{-1, -1}

	// Se for a vez do computador (maximizar)
	if maximizar {
		melhorValor = -1000                        // Define o pior valor inicial
		for linha := 0; linha < tamanho; linha++ { // Itera sobre cada linha do tabuleiro
			for coluna := 0; coluna < tamanho; coluna++ { // Itera sobre cada coluna do tabuleiro
				if tabuleiro[linha][coluna] == vazio { // Verifica se a célula está vazia
					tabuleiro[linha][coluna] = computador                 // Faz um movimento temporário
					valor, _ := minimax(tabuleiro, profundidade+1, false) // Calcula o valor usando minimax recursivamente
					tabuleiro[linha][coluna] = vazio                      // Desfaz o movimento
					if valor > melhorValor {                              // Se o valor atual é melhor, atualiza o melhor valor e movimento
						melhorValor = valor
						melhorMovimento = [2]int{linha, coluna}
					}
				}
			}
		}
	} else { // Se for a vez do jogador (minimizar)
		melhorValor = 1000 // Define o pior valor inicial para o jogador
		for linha := 0; linha < tamanho; linha++ {
			for coluna := 0; coluna < tamanho; coluna++ {
				if tabuleiro[linha][coluna] == vazio {
					tabuleiro[linha][coluna] = jogador                   // Faz um movimento temporário
					valor, _ := minimax(tabuleiro, profundidade+1, true) // Calcula o valor usando minimax recursivamente
					tabuleiro[linha][coluna] = vazio                     // Desfaz o movimento
					if valor < melhorValor {                             // Se o valor atual é melhor, atualiza o melhor valor e movimento
						melhorValor = valor
						melhorMovimento = [2]int{linha, coluna}
					}
				}
			}
		}
	}

	// Retorna o melhor valor encontrado e o movimento correspondente
	return melhorValor, melhorMovimento
}

func verificaFimDoJogo(tabuleiro [tamanho][tamanho]string, jogador string) bool {
	for i := 0; i < tamanho; i++ {
		// verificação linha ou coluna
		if (tabuleiro[i][0] == jogador && tabuleiro[i][1] == jogador && tabuleiro[i][2] == jogador) ||
			(tabuleiro[0][i] == jogador && tabuleiro[1][i] == jogador && tabuleiro[2][i] == jogador) {
			return true
		}
	}
	// verificação primaria e adjacente
	return (tabuleiro[0][0] == jogador && tabuleiro[1][1] == jogador && tabuleiro[2][2] == jogador) ||
		(tabuleiro[0][2] == jogador && tabuleiro[1][1] == jogador && tabuleiro[2][0] == jogador)
}

// verifica se o tabuleiro esta cheio
func empate() bool {
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			if tabuleiro[i][j] == vazio {
				return false
			}
		}
	}
	return true
}
