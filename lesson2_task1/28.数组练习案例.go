package main

import "fmt"

func main() {

	arr2 := [5]int{2, 5, 8, 1, 0}
	a := arrMax(arr2)
	b := arrMin(arr2)
	c := arrSum(arr2)
	d := arrAvg(arr2)

	fmt.Printf("max：%d，min：%d，sum：%d，avg：%0.2f", a, b, c, d)
}

func arrMax(arr [5]int) (_max int) {
	// 假设arr[0]的元素最大
	_max = arr[0]
	for _, val := range arr {
		if val >= _max {
			_max = val
		}
	}
	return
}

func arrMin(arr [5]int) (_min int) {
	// 假设arr[0]的元素最小
	_min = arr[0]
	for _, val := range arr {
		if val <= _min {
			_min = val
		}
	}
	return
}

func arrSum(arr [5]int) (_sum int) {
	for _, val := range arr {
		_sum += val
	}
	return
}

func arrAvg(arr [5]int) (_avg float64) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	_avg = float64(sum / len(arr))
	return
}
