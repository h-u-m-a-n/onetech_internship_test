package quicksort

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		var p int
		p = partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}