package main

import (
	"fmt"
	"igorMSoares/eng-soft-singleton/src"	
)

func main() {
	// Simulando diferentes partes de um sistema usando a mesma fila
	f1 := src.Fila()
	f2 := src.Fila()
	f3 := src.Fila()
	

	fmt.Println("Incluindo \"doc-01\" na fila...")
	f1.AdicionaDocNaFila("doc-01")

	fmt.Println("Incluindo \"doc-02\" na fila...")
	f2.AdicionaDocNaFila("doc-02")

	fmt.Println("Incluindo \"doc-03\" na fila...")
	f3.AdicionaDocNaFila("doc-03")

	f1.ImprimeDocumento()
	f1.ImprimeDocumento()
	f1.ImprimeDocumento()
	f1.ImprimeDocumento()

	fmt.Println("Incluindo \"doc-04\" na fila...")
	f1.AdicionaDocNaFila("doc-04")

	fmt.Println("Incluindo \"doc-05\" na fila...")
	f2.AdicionaDocNaFila("doc-05")

	fmt.Println("Incluindo \"doc-06\" na fila...")
	f3.AdicionaDocNaFila("doc-06")

	f1.RemoveTodosDocumentos()

	fmt.Println("Tentando imprimir após remover todos os documentos...")
	f1.ImprimeDocumento()

	fmt.Println("Tetando remover documento após remover todos os documentos...")
	f2.RemoveDocumento()

	fmt.Println("Tetando remover todos os documentos após remover todos os documentos...")
	f3.RemoveTodosDocumentos()
}
