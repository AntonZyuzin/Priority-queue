package arrheap

import (
	"errors"
	"fmt"
	"time"
)

type BinaryHeap struct {
	heap                []uint // массив с кучей
	insertionComplexity int    // текущая сложность вставки
	rebalanceComplexity int    // текущая сложность перебалансировки дерева
	insertionDuration   time.Duration
	getMaxDuration      time.Duration
	changePriority      time.Duration
}

//вставка
// совсем по тупому - элемент всегда добавляется в конец массива, на котором построено дерево а потом как бы всплывает вверх через сравнения с родителем
func (bh *BinaryHeap) Insert(item int) error {
	start := time.Now()
	bh.insertionComplexity = 0
	elem := uint(item)
	parent := uint((len(bh.heap) - 1) / 2)
	i := uint(len(bh.heap))
	bh.heap = append(bh.heap, elem)
	for parent >= 0 && i > 0 {
		bh.insertionComplexity++
		if bh.heap[i] > bh.heap[parent] {
			temp := bh.heap[i]
			bh.heap[i] = bh.heap[parent]
			bh.heap[parent] = temp
		}
		i = parent
		parent = (i - 1) / 2
	}
	bh.insertionDuration = time.Since(start)
	return nil
}

// это функция перебалансировки. Вызывается каждый раз, когда происходит изъятие максимального элемента
func (bh *BinaryHeap) Heapify(index uint) {
	var temp uint
	left := uint(2*index + 1)
	right := uint(2*index + 2)
	l := uint(len(bh.heap))
	if left < l {
		bh.rebalanceComplexity++
		if bh.heap[index] < bh.heap[left] {
			temp = bh.heap[index]
			bh.heap[index] = bh.heap[left]
			bh.heap[left] = temp
			bh.Heapify(left)
		}
	}
	if right < l {
		bh.rebalanceComplexity++
		if bh.heap[index] < bh.heap[right] {
			temp := bh.heap[index]
			bh.heap[index] = bh.heap[right]
			bh.heap[right] = temp
			bh.Heapify(right)
		}
	}
}

func (bh *BinaryHeap) GetMax() uint {
	start := time.Now()
	x := bh.heap[0]
	bh.heap[0] = bh.heap[len(bh.heap)-1]
	bh.Heapify(0)
	bh.getMaxDuration = time.Since(start)
	return (x)
}

//Очевидно выводит дерево
func (bh *BinaryHeap) PrintTree() {
	i := 0
	k := 1
	l := len(bh.heap)
	for i < l {
		for (i < k) && (i < l) {
			lr := k
			for ; lr > 0; lr-- {
				fmt.Print("-")
			}
			fmt.Print(bh.heap[i])
			fmt.Print(" ")
			i++
		}
		fmt.Println()
		k = k*2 + 1
	}
}

func (bh *BinaryHeap) GetInsertionComplexity() int {
	return bh.insertionComplexity
}

func (bh *BinaryHeap) GetRebalanceComplexity() int {
	return bh.rebalanceComplexity
}

func (bh *BinaryHeap) ClearHeap() {
	bh.heap = make([]uint, 0)
}

func (bh *BinaryHeap) DropRebalanceComplexity() {
	bh.rebalanceComplexity = 0
}

func (bh *BinaryHeap) GetGetMaxDuration() time.Duration {
	return bh.getMaxDuration
}

func (bh *BinaryHeap) GetInsertionDuration() time.Duration {
	return bh.insertionDuration
}

//повышает приоритет на p
func (bh *BinaryHeap) IncreasePrioriy(index, p int) error {
	if index >= 0 && index < len(bh.heap) && p > 0 {
		bh.heap[index] += uint(p)
	} else {
		return errors.New("Что-то пошло не так")
	}
	bh.Heapify(uint(index))
	return nil
}

func (bh *BinaryHeap) ChangePriority(old, new uint) error {
	for k, v := range bh.heap {
		if v == old {
			bh.heap[k] = new
			bh.Heapify(0)
			return nil
		}
	}
	return fmt.Errorf("Элемента cо значением %d не существует", old)
}
