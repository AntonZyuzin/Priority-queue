package main

import (
	"binque/pkg/arrheap/arrheap"
	"binque/pkg/arrheap/linlistqueue"
	"fmt"
	"math/rand"
	"time"
)

const (
	N           = 50000 //количество элементов дерева
	RANGE       = 100   // диапазон от 0 до RANGE в котором рандомно генерятся элементы
	EXPERIMENTS = 10    // количество прогонов для подсчёта средней сложности
)

func main() {
	var compsum float32
	h := &arrheap.BinaryHeap{}
	for e := 0; e < EXPERIMENTS; e++ {
		for i := 0; i < N; i++ {
			rand.Seed(time.Now().UnixNano() + int64(i))
			r := rand.Intn(RANGE) //рандомизация
			h.Insert(r)           //операция вставки
		}
		h.ClearHeap()                                  // Эта функция полностью обнуляет очередь
		compsum += float32(h.GetInsertionComplexity()) // накопление количества сравнений в одной сумме, потом разделим на количество экспериментов
	}

	// кусок кода выше нужен для вычисления сложности вставки, потом куча обнуляется
	for i := 0; i < N; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		//time.Sleep(time.Millisecond * 1)
		r := rand.Intn(RANGE)
		h.Insert(r)
	}

	fmt.Printf("Вставка для текущего дерева заняла %v времени \n", h.GetInsertionDuration())
	fmt.Printf("Взятие максимального элемента и перебалансировка для текущего дерева заняла %v  времени\n", h.GetGetMaxDuration())
	chStart := time.Now()
	err := h.ChangePriority(1, 39)
	chDur := time.Since(chStart)
	fmt.Printf("Время, затраченное на изменение приоритета в куче %v\n", chDur)
	if err != nil {
		fmt.Println(err.Error())
	}
	//	h.PrintTree()
	fmt.Printf("\n\n\n")

	listQueue, errNewQ := linlistqueue.NewQueue()
	if errNewQ != nil {
		fmt.Println(errNewQ.Error())
	}

	// рандом
	insDur := time.Since(time.Now())
	for i := 0; i < N; i++ {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		r := rand.Intn(RANGE)

		insStart := time.Now()
		newHead, errIns := linlistqueue.Insert(listQueue, uint(r))
		insDur = time.Since(insStart)
		listQueue = newHead
		if errIns != nil {
			fmt.Println(errIns.Error())
		}
	}

	fmt.Printf("\n\n\n")
	//	fmt.Println("Была сгенерирована следющая очередь на листе")
	//linlistqueue.PrintQueue(listQueue)

	start := time.Now()
	maxVal, errPop := linlistqueue.PopMax(listQueue)
	duration := time.Since(start)
	if errPop != nil {
		fmt.Println(errPop.Error())
	}
	fmt.Printf("Удаление элемента на листе: %d за %v\n", uint64(maxVal), duration)

	chPStart := time.Now()
	errCh := linlistqueue.ChangePriority(listQueue, 1, 49)
	chPDur := time.Since(chPStart)
	fmt.Printf("Время, затраченное на изменение приоритета в листе %v \n", chPDur)
	if errCh != nil {
		fmt.Println(errCh.Error())
	}
	fmt.Printf("Время, затраченное на вставку в листе %v \n", insDur)
}
