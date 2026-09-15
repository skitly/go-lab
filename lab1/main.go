package main

import (
	"fmt"
	"time"
)
func sumAndDiff(a, b float64) (float64, float64) {
	return a + b, a - b
}
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	// 1. Текущее время и дата
	now := time.Now()
	fmt.Println("Текущая дата и время:", now.Format("2006-01-02 15:04:05"))

	// 2. Переменные разных типов
	var i int = 42
	var f float64 = 3.14
	var s string = "КУ"
	var b bool = true
	fmt.Printf("int: %d, float64: %.2f, string: %s, bool: %t\n", i, f, s, b)

	// 3. Краткая форма объявления
	x := 10
	y := 20
	name := "Я"
	flag := false
	fmt.Println("Краткая форма:", x, y, name, flag)

	// 4. Арифметические операции
	a, c := 15, 4
	fmt.Printf("%d + %d = %d\n", a, c, a+c)
	fmt.Printf("%d - %d = %d\n", a, c, a-c)
	fmt.Printf("%d * %d = %d\n", a, c, a*c)
	fmt.Printf("%d / %d = %d\n", a, c, a/c)
	fmt.Printf("%d %% %d = %d\n", a, c, a%c)

	// 5. Сумма и разность float64
	sum, diff := sumAndDiff(5.5, 2.3)
	fmt.Printf("Сумма: %.2f, Разность: %.2f\n", sum, diff)

	// 6. Среднее значение трёх чисел
	fmt.Printf("Среднее значение: %.2f\n", average(10, 20, 30))
}