package array

import "fmt"

//传递一个数组的指针
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func Array() {
	fmt.Println("————array————")
	//array
	array := [...]float64{7.0, 8.5, 9.1}
	// array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	fmt.Printf("array：%v, x：%v\n", array, x)
	fmt.Println("————————————————")

	fmt.Println("————slice————")
}
