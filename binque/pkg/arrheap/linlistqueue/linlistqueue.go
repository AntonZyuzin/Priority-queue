package linlistqueue

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Next  *Node
	Prev  *Node
	Value uint
}

func NewQueue() (*Node, error) {
	n := &Node{Value: uint(1)}
	return n, nil
}

func Insert(n *Node, val uint) (*Node, error) {
	if n == nil {
		return nil, errors.New("Incoming node ref is nil")
	}
	if val == 0 {
		return n, nil
	}
	newNode := &Node{Next: n, Value: val}
	n.Prev = newNode
	return newNode, nil
}

func PopMax(n *Node) (uint, error) {
	maxNode, err := linearSearchMax(n)
	if err != nil {
		return 0, err
	}
	maxNodeValue := maxNode.Value
	err = deleteNode(maxNode)
	if err != nil {
		return 0, err
	}
	return maxNodeValue, nil
}

//поиск в лоб
func linearSearchMax(n *Node) (*Node, error) {
	if n == nil {
		return nil, errors.New("Argument is nil")
	}
	cur := n
	maxNode := n
	if cur.Next == nil {
		return cur, nil
	}

	for cur.Next != nil {
		if cur.Value > maxNode.Value {
			maxNode = cur
		}
		cur = cur.Next
	}
	return maxNode, nil
}

// удаление перекидывание указателей с предыдущего и последующего друг на друга. Средний элемент теряется и убивается
func deleteNode(n *Node) (err error) {
	err = errors.New("Something went wrong")
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	return nil
}

func PrintQueue(n *Node) error {
	if n == nil {
		return errors.New("Queue is nil")
	}
	cur := n
	fmt.Println()
	fmt.Print(cur.Value)
	for cur.Next != nil {
		cur = cur.Next
		fmt.Print(" ", strconv.FormatUint(uint64(cur.Value), 10))
	}
	fmt.Println()
	return nil
}

func ChangePriority(n *Node, old, new uint) error {
	cur := n
	if cur == nil {
		return fmt.Errorf("Элемента cо значением %d не существует", old)
	}
	if cur.Value == old {
		cur.Value = new
		return nil
	}
	for n.Next != nil {
		if cur.Value == old {
			cur.Value = new
			return nil
		}
		cur = cur.Next
	}
	return nil
}
