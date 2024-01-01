package main

import "fmt"

func main() {
	thisWeekTemp := [7]int{-3, 5, 7}      // [-3 5 7 0 0 0 0]
	rgbColor := [...]uint8{255, 255, 128} // [255 255 128]
	thisWeekTemp2 := [7]int{6: 11, 2: 3}  // [0 0 3 0 0 0 11]
	fmt.Println(thisWeekTemp, rgbColor, thisWeekTemp2)
	//[-3 5 7 0 0 0 0] [255 255 128] [0 0 3 0 0 0 11]

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	// weekTemp [5 4 6 8 11 9 5]
	for _, temp := range weekTemp {
		fmt.Println(temp)
		temp = 0
		fmt.Println(temp)
	}
	// weekTemp [5 4 6 8 11 9 5] ! — значения не изменились
	// если значение элемента не используется, можно опустить вторую переменную
	for i := range weekTemp {
		weekTemp[i] = 0
	}
	// weekTemp [0 0 0 0 0 0 0] ! — значения изменились

	for i, temp := range &weekTemp {
		fmt.Println(i, temp)
	}
}
