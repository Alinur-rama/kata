package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c *Calculator) summa(a, b int) int {
	return a + b
}

func (c *Calculator) raznost(a, b int) int {
	return a - b
}

func (c *Calculator) proizvedenie(a, b int) int {
	return a * b
}

func (c *Calculator) chastnoe(a, b int) int {
	return a / b
}

func (c *Calculator) start(a string) (string, error) {
	b := strings.Fields(a)
	if len(b) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	isRoman := isRimskoe(b[0]) && isRimskoe(b[2])
	if isRimskoe(b[0]) != isRimskoe(b[2]) {
		panic("Выдача паники, так как используются одновременно разные системы счисления")
	}
	var a1, b1 int
	var err error
	if isRoman {
		a1, err = rimskiуNaArab(b[0])
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
		b1, err = rimskiуNaArab(b[2])
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	} else {
		a1, err = strconv.Atoi(b[0])
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
		b1, err = strconv.Atoi(b[2])
		if err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	}
	if a1 < 1 || a1 > 10 || b1 < 1 || b1 > 10 {
		panic("числа должны быть от 1 до 10 включительно")
	}
	var otvet int
	switch b[1] {
	case "+":
		otvet = c.summa(a1, b1)
	case "-":
		otvet = c.raznost(a1, b1)
	case "/":
		if b1 == 0 {
			panic("деление на ноль")
		}
		otvet = c.chastnoe(a1, b1)
	case "*":
		otvet = c.proizvedenie(a1, b1)
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if isRoman {
		if otvet < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел")
		}
		return arabNaRimskiy(otvet), nil
	}
	return strconv.Itoa(otvet), nil
}

func main() {
	c := &Calculator{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	otvet, err := c.start(text)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Ответ:", otvet)

}

func isRimskoe(s string) bool {
	romanNumerals := "IVXLCDM"
	for _, ch := range s {
		if strings.ContainsRune(romanNumerals, ch) {
			return true
		}
	}
	return false
}
func rimskiуNaArab(s string) (int, error) {
	rimskieCif := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	var a, b, c int

	for i := len(s) - 1; i >= 0; i-- {
		c = rimskieCif[rune(s[i])]
		if c >= b {
			a += c
		} else {
			a -= c
		}
		b = c
	}

	return a, nil
}
func arabNaRimskiy(a int) string {
	if a <= 0 {
		return ""
	}

	simvoly := []string{"X", "IX", "V", "IV", "I"}
	cifry := []int{10, 9, 5, 4, 1}

	var s strings.Builder
	for i := 0; i < len(cifry); i++ {
		for a >= cifry[i] {
			s.WriteString(simvoly[i])
			a -= cifry[i]
		}
	}
	return s.String()
}
