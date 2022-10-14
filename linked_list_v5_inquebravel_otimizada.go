package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value Competidor
	Next  *Node
}

type Competidor struct {
	Nome  string
	Vazio string
}

func (l *LinkedList) Append(competidor Competidor) {
	node := &Node{Value: competidor}

	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *LinkedList) showList() {
	node := l.Head
	fmt.Println("\nAqui está a lista: ")
	for node != nil {
		fmt.Println(node.Value.Nome)
		node = node.Next
	}
	fmt.Println(" ")
}

func (l *LinkedList) showPrimeiro() {
	node := l.Head
	fmt.Println(node.Value.Nome)
}

func (l *LinkedList) checarFila(escolha int) {
	if l.Head != nil {
		if escolha == 2 {
			l.removerPrimeiro()
		} else if escolha == 3 {
			l.showList()
		} else if escolha == 4 {
			l.showPrimeiro()
		}
		if escolha == 5 {
			if l.Tail == l.Head {
				fmt.Println("\nNão é possível girar a lista, pois só existe um competidor.")
			} else {
				l.girarFila()
			}
		}
	} else {
		fmt.Println("\nA fila está vazia.")
	}
}
func (l *LinkedList) removerPrimeiro() Competidor {
	if l.Head == nil {
		return Competidor{}
	}
	node := l.Head
	l.Head = node.Next
	l.showList()
	return node.Value

}

func (l *LinkedList) girarFila() {

	node := l.Head
	l.Head = node.Next
	l.Tail.Next = node
	l.Tail = node
	node.Next = nil
	l.showList()
}

func (l *LinkedList) giradorDeFila() {

	node := l.Head
	l.Head = node.Next
	l.Tail.Next = node
	l.Tail = node
	node.Next = nil

}

func (l *LinkedList) girarFilaN(contador int) {
	if l.Tail == l.Head || l.Tail == nil || l.Head == nil {
		fmt.Println("\nNão é possível girar a lista, pois só existe um ou nenhum competidor")
	} else {
		for i := 0; i < contador; i++ {
			l.giradorDeFila()
		}
		l.showList()
	}

}

func main() {

	list := LinkedList{}
	var escolha = 0

	for escolha != 7 {
		fmt.Println("\nEscolha uma das opções abaixo: ")
		fmt.Println("1 - Adicionar competidor\n2 - Remover competidor\n3 - Mostrar lista\n4 - Mostrar o primeiro\n5 - Girar lista\n6 - Girar lista N vezes\n7 - Finalizar")
		fmt.Scanf("%d\n", &escolha)
		if escolha == 1 {
			var comp string
			fmt.Println("\nInforme o nome do competidor: ")
			fmt.Scanf("%s\n", &comp)
			fmt.Println("")
			list.Append(Competidor{Nome: comp})

		} else if escolha == 2 {
			list.checarFila(escolha)
		} else if escolha == 3 {
			list.checarFila(escolha)
		} else if escolha == 4 {
			list.checarFila(escolha)
		} else if escolha == 5 {
			list.checarFila(escolha)
		} else if escolha == 6 {
			contador := 0
			fmt.Println("Informe quantas vezes deseja girar a fila: ")
			fmt.Scanf("%d\n", &contador)
			list.girarFilaN(contador)
		} else if escolha > 7 || escolha <= 0 {
			fmt.Println("Opção inválida.")
		}

	}

	if escolha == 7 {
		fmt.Println("Finalizado.")
	} else {
		fmt.Println("Opção inválida.")
	}
}
