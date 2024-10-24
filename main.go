#Задание 1.1
package main

import "fmt"

func addNumberDigits(number int32) int32 {
	//Максимальное количество цифр в числе типа int32 - 10, 
	//поэтому повторяем операцию 10 раз, чтобы не использовать цикл
  digit1 := num % 10
	digit2 := (num / 10) % 10
	digit3 := (num / 100) % 10
	digit4 := (num / 1000) % 10
  digit5 := (num / 10000) % 10
  digit6 := (num / 100000) % 10
  digit7 := (num / 1000000) % 10
  digit8 := (num / 10000000) % 10
  digit9 := (num / 100000000) % 10
  digit10 := (num / 100000000) % 10

	return digit1 + digit2 + digit3 + digit4 + digit5 + digit6 + digit7 + digit8 + digit9 + digit10
}

func main() {
	fmt.Println(addNumberDigits(1234567890))
}
