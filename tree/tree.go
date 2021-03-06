package tree

import (
	"fmt"
	"strings"
)

// Creo un tipo personalizzato a cui assegno valori costanti
type state int

const (
	empty state = 0
	half  state = 1
	full  state = 2
)

// Node identify the single element of a tree
// value: key of the node
// left & right: child of Node
// parent: parent of Node
// fill: indicates the status of Node (empty = no child is full, half = one child is full, full = cannot have any new child)
type Node struct {
	value               string
	left, right, parent *Node
	fill                state
}

// ReadPreOrderTree gets a string written as PreOrder visit and returns a its binary tree
func ReadPreOrderTree(elements []string) *Node {
	var tree *Node

	// Populate the tree
	for _, el := range elements {
		tree = add(tree, el)
	}
	return tree
}

// Sort puts the keys of the tree in a slice as in order visit,
// if print == true, it also prints the slice on file
func Sort(tree *Node, print bool) []string {
	var elements []string
	// saves the elements in a slice
	elements = appendValues(elements, tree)
	// print, avoid NULL
	if print == true {
		for _, val := range elements {
			if !strings.Contains(val, "NULL") {
				fmt.Printf("%s ", val)
			}
		}
	}
	return elements
}

// add adds the node to the tree
func add(t *Node, value string) *Node {
	// Se l'albero è vuoto, lo inizializzo e termino
	if t == nil {
		t = new(Node)
		t.value = value
		if strings.Contains(value, "NULL") {
			t.fill = full
		} else {
			t.fill = empty
		}
		setState(t)
		return t
	}
	// Se il riempimento del nodo è empty...
	if t.fill == empty {
		// ...vado in ricorsione a sx
		t.left = add(t.left, value)
		t.left.parent = t
		// Altrimenti se è half...
	} else if t.fill == half {
		// ...vado in ricorsione a dx
		t.right = add(t.right, value)
		t.right.parent = t
		// Altrimenti se pieno...
	} else if t.fill == full {
		// ...vado al ramo destro del genitore
		t.parent.right = add(t.parent.right, value)
		t.parent.right.parent = t
	}
	// Aggiorno lo stato del nodo e termino
	setState(t)
	return t
}

func setState(t *Node) {
	// Se il valore del nodo è NULL, chiude
	if strings.Contains(t.value, "NULL") {
		return
	}
	// Se sx non esiste...
	if t.left == nil {
		// ...imposta come empty
		t.fill = empty
		// altrimenti, se a sinistra trova NULL o è impostato come full...
	} else if strings.Contains(t.left.value, "NULL") || t.left.fill == full {
		// ...imposta come half
		t.fill = half
	}
	// se dx esiste e ha stato pieno...
	if t.right != nil && t.right.fill == full {
		// ... imposta il nodo come full
		t.fill = full
	} else {
		// vuoto altrimenti
		t.fill = empty
	}
}

// appendValues puts the elements on a slice
func appendValues(values []string, t *Node) []string {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
