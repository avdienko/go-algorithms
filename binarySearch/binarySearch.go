package binarysearch

func binarySearch(arrayzor []int, toFind int) int {

	var low, high int
	low = arrayzor[0]
	high = arrayzor[len(arrayzor)-1]

	if toFind > high || toFind < low {
		return -1
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case toFind < arrayzor[mid]:
			high = mid - 1

		case toFind > arrayzor[mid]:
			low = mid + 1

		case toFind == arrayzor[mid]:
			return mid
		}
	}
	return -1
}
