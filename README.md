# Singleton Design Pattern

Exercício para a aula de Engenharia de Software III - FATEC Baixada Santista

## Objetivo

Criar uma classe Fila que implementa o design pattern Singleton, ou seja, só será permitido existir uma única instância da classe Fila.

A classe Fila deverá expôr os seguintes métodos:

- Fila(): retorna a instância de Fila
- ImprimeDocumento(): simula a impressão de um documento, removendo-o da fila (em uma fila o primeiro elemento a ser adicionado será o primeiro elemento a ser impresso)
- RemoveDocumento(): remove o último documento da fila
- RemoveTodosDocumentos(): remove todos os documentos da fila
- AdicionaDocNaFila(nomeDoDocumento): recebe uma string simbolizando o nome do documento e inclui o nome do documento na fila

## Metodologia

A classe Fila foi implementada utilizando a linguagem [Go](https://go.dev/).
Apesar de Go não ser estritamente orientada a objetos, a implementação de singleton foi simulada utilizando o conceito de packages.

### Packages

O arquivo principal (`main.go`) está no package chamado `main`.
O arquivo que implementa a "classe" Fila (`Fila.go`) está no package chamado `src`.

Em Go, para exportar qualquer coisa para fora de um package, o método ou variável deve ser ser nomeado com letra maiúscula.
Com isso em mente, o arquivo `Fila.go` declara uma struct `fila` que armazena a fila de impressão (nome dos documentos incluídos na fila).
A struct `fila` não está sendo exportada (nome em minúsculo), dessa forma nenhum outro package poderá criar structs do tipo `fila`.
O arquivo `Fila.go` define uma variável `queue` do tipo ponteiro para `fila`. Inicialmente essa variável é apenas declarada, mas não aponta pra nenhuma struct do tipo `fila`, simulando um atributo privado que guarda a referência para a instância de `fila`.

O arquivo `Fila.go` exporta apenas os métodos "públicos", que serão importados e utilizados no arquivo `main.go`

## Executar

Para executar o programa, rode o comando `go run main.go`.
Caso não tenha a linguagem Go instalada, poderá utilizar a versão compilada, executando diretamente o binário `main`.
