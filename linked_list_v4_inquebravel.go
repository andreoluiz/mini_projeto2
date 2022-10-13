package main

import (
	"fmt"
)

type Linked_list struct {
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

func (l *Linked_list) Append(competidor Competidor) {
	node := &Node{Value: competidor}

	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *Linked_list) show_List() {
	node := l.Head
	fmt.Println("\nAqui está a lista: ")
	for node != nil {
		fmt.Println(node.Value.Nome)
		node = node.Next
	}
	fmt.Println(" ")
}

func (l *Linked_list) show_Primeiro() {
	node := l.Head
	fmt.Println(node.Value.Nome)
}

func (l *Linked_list) checar_Fila(escolha int) {
	if l.Head != nil {
		if escolha == 2 {
			l.remover_Primeiro()
		} else if escolha == 3 {
			l.show_List()
		} else if escolha == 4 {
			l.show_Primeiro()
		}
		if escolha == 5 {
			if l.Tail == l.Head {
				fmt.Println("\nNão é possível girar a lista, pois só existe um competidor.")
			} else {
				l.girar_Fila()
			}
		}
	} else {
		fmt.Println("\nA fila está vazia.")
	}
}
func (l *Linked_list) remover_Primeiro() Competidor {
	if l.Head == nil {
		return Competidor{}
	}
	node := l.Head
	l.Head = node.Next
	l.show_List()
	return node.Value

}

func (l *Linked_list) girar_Fila() {

	node := l.Head
	l.Head = node.Next
	l.Tail.Next = node
	l.Tail = node
	node.Next = nil
	l.show_List()
}

func (l *Linked_list) girador_de_Fila() {

	node := l.Head
	l.Head = node.Next
	l.Tail.Next = node
	l.Tail = node
	node.Next = nil

}

func (l *Linked_list) girar_Fila_N(contador int) {
	if l.Tail == l.Head || l.Tail == nil || l.Head == nil {
		fmt.Println("\nNão é possível girar a lista, pois só existe um ou nenhum competidor")
	} else {
		for i := 0; i < contador; i++ {
			l.girador_de_Fila()
		}
		l.show_List()
	}

}

func main() {

	list := Linked_list{}
	var escolha int

	fmt.Println("Escolha uma das opções abaixo: ")
	fmt.Println("1 - Adicionar competidor\n2 - Remover competidor\n3 - Mostrar lista\n4 - Mostrar o primeiro\n5 - Girar lista\n6 - Girar lista N vezes\n7 - Finalizar")
	fmt.Scanf("%d\n", &escolha)

	if escolha == 1 {
		var comp string
		fmt.Println("\nInforme o nome do competidor: ")
		fmt.Scanf("%s\n", &comp)
		list.Append(Competidor{Nome: comp})
	} else if escolha == 2 {
		list.checar_Fila(escolha)
	} else if escolha == 3 {
		list.checar_Fila(escolha)
	} else if escolha == 4 {
		list.checar_Fila(escolha)
	} else if escolha == 5 {
		list.checar_Fila(escolha)
	} else if escolha == 6 {
		contador := 0
		fmt.Println("Informe quantas vezes deseja girar a fila: ")
		fmt.Scanf("%d\n", &contador)
		list.girar_Fila_N(contador)
	}

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
			list.checar_Fila(escolha)
		} else if escolha == 3 {
			list.checar_Fila(escolha)
		} else if escolha == 4 {
			list.checar_Fila(escolha)
		} else if escolha == 5 {
			list.checar_Fila(escolha)
		} else if escolha == 6 {
			contador := 0
			fmt.Println("Informe quantas vezes deseja girar a fila: ")
			fmt.Scanf("%d\n", &contador)
			list.girar_Fila_N(contador)
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
