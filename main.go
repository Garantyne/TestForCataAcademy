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
	num1, num2 := decipheringNumerals(arifmetic)
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
func decipheringNumerals(arifmet string) (int, int) {
	arifmet = strings.TrimSpace(arifmet)
	signIndex := strings.Index(arifmet, getSign(arifmet))
	if signIndex < 1 {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}
	str1 := arifmet[:signIndex]
	str1 = strings.TrimSpace(str1)
	str2 := arifmet[signIndex+1:]
	str2 = strings.TrimSpace(str2)

	num1, isRomeSystem1, isArabSystem1 := parseDigit(str1)
	num2, isRomeSystem2, isArabSystem2 := parseDigit(str2)
	if isRomeSystem1 && isArabSystem2 || isRomeSystem2 && isArabSystem1 {
		panic("Разные системы счисления, аварийное завершение")
	}
	if (isRomeSystem2 && isRomeSystem1) && num1 < num2 && getSign(arifmet) == "-" {
		panic("Аварийное завершение в римской системе небыло отрицательных чисел")
	}

	if isRomeSystem1 && isRomeSystem2 {
		showRomeDigitAnswer(num1, num2, arifmet)
	}

	return num1, num2
}

func showRomeDigitAnswer(num1 int, num2 int, arifmet string) {
	result := calculator(num1, num2, getSign(arifmet))
	if result == 1 {
		fmt.Println("I")

	} else if result == 2 {
		fmt.Println("II")
	} else if result == 3 {
		fmt.Println("III")
	} else if result == 4 {
		fmt.Println("IV")
	} else if result == 5 {
		fmt.Println("V")
	} else if result == 6 {
		fmt.Println("VI")
	} else if result == 7 {
		fmt.Println("VII")
	} else if result == 8 {
		fmt.Println("VIII")
	} else if result == 9 {
		fmt.Println("IX")
	} else if result == 10 {
		fmt.Println("IX")
	} else if result == 11 {
		fmt.Println("XI")
	} else if result == 12 {
		fmt.Println("XII")
	} else if result == 13 {
		fmt.Println("XII")
	} else if result == 14 {
		fmt.Println("XIV")
	} else if result == 15 {
		fmt.Println("XV")
	} else if result == 16 {
		fmt.Println("XVI")
	} else if result == 17 {
		fmt.Println("XVII")
	} else if result == 18 {
		fmt.Println("XVIII")
	} else if result == 19 {
		fmt.Println("XIX")
	} else if result == 20 {
		fmt.Println("XX")
	} else {
		fmt.Println("Некорректное число")
	}
}

func parseDigit(digit string) (int, bool, bool) {
	var num1 int
	var isRomeSystem bool = false
	var isArabSystem bool = false
	if digit == "I" {
		num1 = 1
		isRomeSystem = true
	} else if digit == "II" {
		num1 = 2
		isRomeSystem = true
	} else if digit == "III" {
		num1 = 3
		isRomeSystem = true
	} else if digit == "IV" {
		num1 = 4
		isRomeSystem = true
	} else if digit == "V" {
		num1 = 5
		isRomeSystem = true
	} else if digit == "VI" {
		num1 = 6
		isRomeSystem = true
	} else if digit == "VII" {
		num1 = 7
		isRomeSystem = true
	} else if digit == "VIII" {
		num1 = 8
		isRomeSystem = true
	} else if digit == "IX" {
		num1 = 9
		isRomeSystem = true
	} else if digit == "X" {
		num1 = 10
		isRomeSystem = true
	} else {
		var err error
		num1, err = strconv.Atoi(digit)
		isArabSystem = true
		if err != nil {
			fmt.Println("Вывод ошибки, так как формат математической "+
				"операции не удовлетворяет заданию — два операнда и один "+
				"оператор (+, -, /, *).\n\n", err)
			panic("Ошибка")
		}
	}
	return num1, isRomeSystem, isArabSystem
}
