package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sign — определяет знак числа
func sign(n int) string {
	switch {
	case n > 0:
		return "Positive"
	case n < 0:
		return "Negative"
	default:
		return "Zero"
	}
}

// strLength — длина строки в рунах (корректно для кириллицы)
func strLength(s string) int {
	return len([]rune(s))
}

// Rectangle — прямоугольник
type Rectangle struct {
	Width, Height float64
}

// Area — площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// avgInt — среднее двух целых чисел
func avgInt(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Чётное или нечётное
	fmt.Print("Введите целое число: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка: введите целое число")
		return
	}
	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
	fmt.Println("Знак числа:", sign(n))

	// 2. Числа от 1 до 10
	fmt.Print("Числа от 1 до 10: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 3. Длина строки
	fmt.Println("Длина строки \"Привет, мир!\":", strLength("Привет, мир!"))

	// 4. Площадь прямоугольника
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Площадь прямоугольника: %.2f\n", rect.Area())

	// 5. Среднее двух целых
	fmt.Printf("Среднее 4 и 7: %.2f\n", avgInt(4, 7))
}