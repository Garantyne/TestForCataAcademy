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
	var num3 int
	var str string
	num3 += result % 10
	result /= 10
	if result < 5 {
		for i := 0; i < result; i++ {
			str += "X"
		}
	} else if result < 9 {
		str = "L"
		for i := 5; i < result; i++ {
			str += "X"
		}
	} else if result == 10 {
		str = "LL"
	}
	if num3 == 1 {
		str += "I"
	} else if num3 == 2 {
		str += "II"
	} else if num3 == 3 {
		str += "III"
	} else if num3 == 4 {
		str += "IV"
	} else if num3 == 5 {
		str += "V"
	} else if num3 == 6 {
		str += "VI"
	} else if num3 == 7 {
		str += "VII"
	} else if num3 == 8 {
		str += "VIII"
	} else if num3 == 9 {
		str += "IX"
	}
	fmt.Println(str)
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
		if num1 < 0 || num1 > 10 {
			panic("Вы ввели некорректное число ")
		}
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
