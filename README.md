# Jogo da Velha em Golang

Este projeto é uma implementação de um jogo da velha (Tic-Tac-Toe) em Golang, onde o jogador enfrenta um agente inteligente que utiliza o algoritmo Minimax para responder às jogadas com o objetivo de ganhar ou, no mínimo, empatar.

## Descrição

- **Algoritmo Minimax**: O jogo utiliza o algoritmo Minimax, que é uma técnica clássica em inteligência artificial para jogos de dois jogadores. O algoritmo avalia todas as possíveis jogadas e escolhe a que maximiza as chances de vitória do computador, minimizando as chances de vitória do oponente.

- **Validação de Entrada**: Foram implementadas várias verificações para garantir que as entradas do jogador sejam válidas, ou seja, que sejam números dentro do intervalo permitido e que correspondam a posições livres na matriz.

- **Gabarito para o Jogador**: Para facilitar a jogabilidade, um gabarito é exibido, mostrando as posições numéricas correspondentes na matriz, ajudando o jogador a escolher sua jogada.

## Como Jogar

1. **Clonar o repositório**:
   ```bash
   git clone https://github.com/AdrilesTheCreator/jogo-da-velha-golang.git
   ```
2. **Navegar até o diretório do projeto**:
   ```bash
   cd jogo-da-velha-golang
   ```
3. **Executar o jogo**:
   ```bash
   go run main.go
   ```
4. **Instruções**: Siga as instruções na tela para inserir suas jogadas. Digite o número correspondente à posição desejada conforme o gabarito.

## Exemplo de Uso

Ao iniciar o jogo, será exibida uma matriz 3x3 numerada. O jogador deve escolher uma posição digitando o número correspondente. O computador então faz sua jogada com base no algoritmo Minimax, e o jogo continua até que haja um vencedor ou empate.

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Algoritmo**: Minimax para tomada de decisões

## Contribuições

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, fique à vontade para enviar um pull request.
