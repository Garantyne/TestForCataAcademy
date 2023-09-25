package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите арифметическое выражение")
	arifmetic, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(arifmetic)
	/*num1, num2 := getDigit(arifmetic)
	fmt.Println(getSign(arifmetic))
	fmt.Println(calculator(num1, num2, getSign(arifmetic)))*/
	num1, num2 := decipheringRomanNumerals(arifmetic)
	fmt.Println(calculator(num1, num2, getSign(arifmetic)))

}
func getDigit(arifmet string) (int, int) {
	//удаляем все пробелы если вдруг их ввели
	arifmet = strings.TrimSpace(arifmet)
	signIndex := strings.Index(arifmet, getSign(arifmet))
	str1 := arifmet[:signIndex]
	str1 = strings.TrimSpace(str1)
	str2 := arifmet[signIndex+1:]
	str2 = strings.TrimSpace(str2)
	//подготавливаем числа которыре будем извлекать из строк
	var number1 int
	var number2 int
	if signIndex != -1 {
		var err error
		number1, err = strconv.Atoi(str1)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			return number1, number2
		}
		number2, err = strconv.Atoi(str2)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			return number1, number2
		}
	}
	return number1, number2
}

// получаем знак
func getSign(arifmet string) string {
	signIndex := strings.Index(arifmet, " ")
	sign := arifmet[signIndex+1 : signIndex+2]
	sign = strings.TrimSpace(sign)
	return sign
}

// производим арифметическую операцию
func calculator(num1, num2 int, sign string) int {
	var result int
	switch sign {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}

// расшифровка римских цифр
func decipheringRomanNumerals(arifmet string) (int, int) {
	arifmet = strings.TrimSpace(arifmet)
	signIndex := strings.Index(arifmet, getSign(arifmet))
	str1 := arifmet[:signIndex]
	str1 = strings.TrimSpace(str1)
	str2 := arifmet[signIndex+1:]
	str2 = strings.TrimSpace(str2)

	num1 := parseDigit(str1)
	num2 := parseDigit(str2)

	return num1, num2
}

func parseDigit(digit string) int {
	var num1 int
	if digit == "I" {
		num1 = 1
	} else if digit == "II" {
		num1 = 2
	} else if digit == "III" {
		num1 = 3
	} else if digit == "IV" {
		num1 = 4
	} else if digit == "V" {
		num1 = 5
	} else if digit == "VI" {
		num1 = 6
	} else if digit == "VII" {
		num1 = 7
	} else if digit == "VIII" {
		num1 = 8
	} else if digit == "IX" {
		num1 = 9
	} else if digit == "X" {
		num1 = 10
	} else {
		var err error
		num1, err = strconv.Atoi(digit)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			panic("Ошибка")
		}
	}
	return num1
}
