# Desafio Go: Ping Pong (Concorrência e Sincronização) 🏓

![Banner Go](https://assets.dio.me/Ko22-i7zgPJFpuJaQHdgZMyEIiOA8bzSOqqFKv8Hj5I/f:webp/q:80/L2FydGljbGVzL2NvdmVyLzkzNDAwNjEzLTllNGMtNDRjOS1iNDJhLTlhMjk3MmFmMzA2MC5wbmc)

Este projeto foi desenvolvido para exercitar conceitos avançados de concorrência e comunicação entre processos na linguagem **Go**. O algoritmo utiliza as ferramentas nativas da linguagem para coordenar duas tarefas independentes que precisam trabalhar em harmonia.

## 📝 Descrição do Desafio

O objetivo é criar um sistema de trocas de mensagens sincronizadas entre duas funções, simulando uma partida de Ping Pong com as seguintes regras:

* Uma goroutine deve ser responsável por exibir a palavra **"ping"**.
* Outra goroutine deve ser responsável por exibir a palavra **"pong"**.
* As palavras devem aparecer no terminal de forma **alternada** (um ping sempre seguido de um pong).
* O programa deve encerrar automaticamente após atingir um limite pré-definido de rodadas.

## 🛠️ Tecnologias Utilizadas

* **Go (Golang)**: Linguagem focada em alta performance e concorrência nativa.
* **Pacote fmt**: Utilizado para a exibição dos resultados no console.
* **Pacote sync**: Utilizado para o gerenciamento e espera das rotinas de execução.

## 🚀 Como Rodar o Código

1. Certifique-se de ter o Go instalado em sua máquina.
2. Copie o código para um arquivo chamado `main.go`.
3. Execute o comando no seu terminal:

```bash
go run main.go

```

## 🧠 Conceitos Aplicados

* **Goroutines**: Execução de funções de forma assíncrona e concorrente.
* **Canais (`chan`)**: Utilizados como conduítes de comunicação e sincronização para evitar condições de corrida.
* **WaitGroups (`sync.WaitGroup`)**: Mecanismo para coordenar o encerramento do programa principal, garantindo que ele espere as goroutines terminarem.
* **Escopo e Constantes**: Uso de constantes em nível de pacote para centralizar a configuração do limite de jogadas.
* **Defer**: Utilizado para garantir que a sinalização de término (`Done`) seja executada ao final do ciclo das funções.