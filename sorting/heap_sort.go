package main

import "fmt"

func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
}

func heapify(arr []int, i, size int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < size && arr[left] > arr[largest] {
		largest = left
	}

	if right < size && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, size)
	}
}

func heapSort(arr []int) {
	buildMaxHeap(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func main() {
	arr := []int{9, 2, 5, 3, 8, 1, 6}
	heapSort(arr)
	fmt.Println(arr)
}
