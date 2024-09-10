package src

import "fmt"

type fila struct {
	queue []string
}

var queue *fila

func Fila() *fila {
	// Verifica se queue já foi instanciado anteriormente
	if queue == nil {
		// Cria uma instância de fila
		queue = &fila{}
	}

	return queue
}

func (f *fila) ehFilaVazia() bool {
	if len(f.queue) == 0 {
		fmt.Println(">> A fila de impressão está vazia.")
		fmt.Println()
		return true
	}

	return false
}

func (f *fila) AdicionaDocNaFila(doc string) {
	f.queue = append(f.queue, doc)

	fmt.Println(">> Documento adicionado na fila de impressão:", doc)
	fmt.Println()
}

func (f *fila) ImprimeDocumento() {
	if f.ehFilaVazia() {
		return
	}

	doc := f.queue[0]
	f.queue = f.queue[1:]

	fmt.Printf("Imprindo documento \"%v\"\n\n", doc)
}

func (f *fila) RemoveDocumento() {
	if f.ehFilaVazia() {
		return
	}

	fmt.Println("Removendo documento da fila...")

	doc := f.queue[len(f.queue)-1]
	f.queue = f.queue[:len(f.queue)-1]

	fmt.Printf("Documento \"%v\" removido com sucesso!\n\n", doc)
}

func (f *fila) RemoveTodosDocumentos() {
	if f.ehFilaVazia() {
		return
	}

	fmt.Printf("Removendo todos os %v documentos...\n", len(f.queue))
	f.queue = []string{}
	fmt.Println("Documentos removidos com sucesso! A fila está vazia.")
	fmt.Println()
}
