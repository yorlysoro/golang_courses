package middlelinkedlist

import (
	"fmt"
	"strings"
)

// ListNode es la definición de nodo para una lista enlazada simple.
type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode devuelve el nodo medio de la lista enlazada.
// Comportamiento: si la lista tiene longitud par, devuelve el segundo de los dos nodos medios.
// Estrategia: dos punteros (slow, fast). Fast avanza 2 pasos por iteración y slow 1 paso.
// Cuando fast llega al final, slow está en el medio.
// Tiempo: O(n). Espacio: O(1).
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// BuildList construye una lista enlazada a partir de una slice de enteros.
// Devuelve la cabeza de la lista (nil si vals está vacío).
func BuildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

// ListToString convierte la lista a una representación string como "1 -> 2 -> 3".
func ListToString(head *ListNode) string {
	if head == nil {
		return "nil"
	}
	parts := make([]string, 0)
	for cur := head; cur != nil; cur = cur.Next {
		parts = append(parts, fmt.Sprintf("%d", cur.Val))
	}
	return strings.Join(parts, " -> ")
}

/*func main() {
	examples := [][]int{
		{1, 2, 3, 4, 5},      // longitud impar -> medio = 3
		{1, 2, 3, 4, 5, 6},   // longitud par -> medio esperado = 4 (segundo medio)
		{1},                  // un solo nodo
		{},                   // lista vacía
		{10, 20, 30, 40, 50}, // otro ejemplo
	}

	for _, vals := range examples {
		head := BuildList(vals)
		fmt.Printf("Lista: %s\n", ListToString(head))
		mid := middleNode(head)
		if mid == nil {
			fmt.Println("Nodo medio: nil\n")
		} else {
			fmt.Printf("Nodo medio: %d\n\n", mid.Val)
		}
	}
}*/
