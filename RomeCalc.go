package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
