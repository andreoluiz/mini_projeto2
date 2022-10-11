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
	Nome      string
	Sobrenome string
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

func (l *Linked_list) show_Ultimo() {
	node := l.Tail
	fmt.Println(node.Value.Nome)
}

func (l *Linked_list) remover_Primeiro() Competidor {
	node := l.Head
	l.Head = node.Next
  return node.Value
}

func (l *Linked_list) girar_Fila() {
	node := l.Head
	l.Head = node.Next
	l.Tail.Next = node
	l.Tail = node
	node.Next = nil
  //lista := Linked_list{}
  //salvar := l.remover_Primeiro()
  //l.Append(salvar)
}

func (l *Linked_list) girar_Fila_N(n int) {
	for i := 0; i < n; i++ {
		l.girar_Fila()
	}
}

func main() {

	list := Linked_list{}

	gabriel := Competidor{Nome: "Gabriel", Sobrenome: "Safado"}
	andre := Competidor{Nome: "Andre", Sobrenome: "FiDoGo"}
	denis := Competidor{Nome: "Denis", Sobrenome: "Uiulson"}
	isaque := Competidor{Nome: "Isaque", Sobrenome: "Viçosa"}
	vitor := Competidor{Nome: "Vitor", Sobrenome: "oMelo"}
  joao := Competidor{Nome: "Joao", Sobrenome: "ABC"}

	list.Append(gabriel)
	list.Append(andre)
	list.Append(denis)
	list.Append(isaque)
	list.Append(vitor)
  list.Append(joao)

  list.show_List()
  list.girar_Fila_N(5)
	list.show_List()
  
}