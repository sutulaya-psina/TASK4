###Задание 1.1

package main

import (
	"fmt"
)

func findDigitSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func main() {
	fmt.Println(findDigitSum(1234))
}

###Задание 1.2

package main

import (
	"fmt"
  "strconv"
	"strings"
)

func ToF(temp int) float32 {
	return float32(temp)* 9.0 / 5.0 + 32
}

func ToC(temp int) float32 {
	return (float32)(temp-32) * 5.0 / 9.0
}


func main() {
	cel := "25 (Celsius)"
	temp1, _ := strconv.Atoi(strings.Split(cel, " ")[0])

	fmt.Printf("%f (Farhrenheit)\n", ToF(temp1))

	far := "32 (Farhrenheit)"
	temp2, _ := strconv.Atoi(strings.Split(far, " ")[0])

	fmt.Printf("%f (Celsius)\n", ToC(temp2))

}

###Задание 1.3

package main

import ("fmt"
)

func doubleArray(input []int) []int {
	output := make([]int, len(input))
	for i, v := range input {
	 output[i] = v * 2
	}
	return output
   }

   func main() {
		arr := []int{1, 2, 3, 4}
		result := doubleArray(arr)

   fmt.Printl(, arr)
   fmt.Printl(, result)
   }

###Задание 1.4

package main

import (
 "fmt"
 "os"
 "strings"
)

func main() {
 args := os.Args[1:]
 if len(args) == 0 {
  fmt.Println("Осторожно,", "горячий кофе")
  os.Exit(1)
 }

 result := strings.Join(args, " ")
 fmt.Println(result)
}

###Задание 1.5

package main

import (
 "fmt"
 "math"
)

type Point struct {
 X float64
 Y float64
}

func Distance(p1, p2 Point) float64 {
 dX := p2.X - p1.X
 dY := p2.Y - p1.Y
 return math.Sqrt(dX*dX + dY*dY)
}

func main() {
 p1 := Point{X: 1, Y: 1}
 p2 := Point{X: 4, Y: 5}

 distance := Distance(p1, p2)
 fmt.Printf("Расстояние между (%f, %f) и (%f, %f) равно %.1f\n",
  p1.X, p1.Y, p2.X, p2.Y, distance)
}

###Задание 2.1

package main

import (
 "fmt"
)

func main() {
 var number int
 fmt.Print("Введите число: ")
 fmt.Scan(&number)

 if number%2 == 0 {
  fmt.Println("Четное")
 } else {
  fmt.Println("Нечетное")
 }
}

###Задание 2.2

package main

import (
 "fmt"
)

func isLeapYear(year int) bool {
 return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
 var year int
 fmt.Print("Введите год: ")
 fmt.Scan(&year)

 if isLeapYear(year) {
  fmt.Printf("%d - Високосный\n", year)
 } else {
  fmt.Printf("%d - Невисокосный\n", year)
 }
}

###Задание 2.3

package main

import (
 "fmt"
)

func main() {
 var a, b, c int

 fmt.Print("Введите первое число: ")
 fmt.Scanln(&a)

 fmt.Print("Введите второе число: ")
 fmt.Scanln(&b)

 fmt.Print("Введите третье число: ")
 fmt.Scanln(&c)

 max := a
 if b > max {
  max = b
 }
 if c > max {
  max = c
 }

 fmt.Printf("Наибольшее число: %d\n", max)
}

###Задание 2.4

package main

import (
 "fmt"
)

func main() {
 var age int
 fmt.Print("Введите ваш возраст: ")
 fmt.Scan(&age)

 var group string
 if age <= 12 {
  group = "Ребенок"
 } else if age > 12 && age <= 18 {
  group = "Подросток"
 } else if age > 18 && age <= 60 {
  group = "Взрослый"
 } else {
  group = "Пожилой"
 }

 fmt.Printf("Вы %s\n", group)
}

###Задание 2.5

package main

import (
 "fmt"
)

func main() {
 var number int
 fmt.Print("Введите число: ")
 fmt.Scan(&number)

 if number%3 == 0 && number%5 == 0 {
  fmt.Println("Делится")
 } else {
  fmt.Println("Не делится")
 }
}


###Задание 3.1

package main

import (
 "fmt"
)

func main() {
 var n int
 fmt.Print("Введите число: ")
 fmt.Scan(&n)

 factorial := 1
 for i := 2; i <= n; i++ {
  factorial *= i
 }

 fmt.Printf("%d! = %d\n", n, factorial)
}

###Задание 3.2

package main

import (
 "fmt"
)

func main() {

 var n int
 fmt.Print("Введите количество чисел Фибоначчи: ")
 fmt.Scanln(&n)

 a, b := 0, 1
 for i := 0; i < n; i++ {
  fmt.Print(a, " ")
  a, b = b, a+b
 }
}

###Задание 3.3

package main

import (
 "fmt"
)

func reverseArray(arr []int) []int {
 for i := 0; i < len(arr)/2; i++ {
  arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
 }
 return arr
}

func main() {
 inputArr := []int{1, 2, 3, 4, 5}

 reversedArr := reverseArray(inputArr)

 fmt.Println(reversedArr)
}

###Задание 3.4

package main

import (
 "fmt"
)

func sieveOfEratosthenes(limit int) []int {
 isPrime := make([]bool, limit+1)
 for i := range isPrime {
  isPrime[i] = true
 }

 isPrime[0] = false
 isPrime[1] = false

 for i := 2; i*i <= limit; i++ {
  if isPrime[i] {
   for j := i * i; j <= limit; j += i {
    isPrime[j] = false
   }
  }
 }

 primes := []int{}
 for i := 2; i <= limit; i++ {
  if isPrime[i] {
   primes = append(primes, i)
  }
 }

 return primes
}

func main() {
 var n int
 fmt.Print("Введите предел: ")
 fmt.Scan(&n)

 primes := sieveOfEratosthenes(n)
 fmt.Println("Простые числа до", n, ":")
 fmt.Println(primes)
}

###Задание 3.5

package main

import (
 "fmt"
)

func sumArray(arr []int) int {
 sum := 0
 for _, value := range arr {
  sum += value
 }
 return sum
}

func main() {
 inputArr := []int{1, 2, 3, 4, 5}
 result := sumArray(inputArr)
 fmt.Println(result)
}


